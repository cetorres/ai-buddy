package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/models"
	"github.com/cetorres/ai-buddy/pattern"
	"github.com/cetorres/ai-buddy/util"
)

const DEBUG = false

//go:embed static/*
var	staticDir embed.FS

type Page struct {
	Title string
	Body string
}

func CreateHTTPServer() {
	// Routes
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/settings", handleSettingsPage)
	http.HandleFunc("/settings_values", handleSettingsValue)
	http.HandleFunc("POST /save_settings", handleSaveSettings)
	http.HandleFunc("/providers", handleGetProviders)
	http.HandleFunc("/models", handleGetModels)
	http.HandleFunc("/patterns", handleGetPatterns)
	http.HandleFunc("/version", handleGetVersion)
	http.HandleFunc("POST /execute", handleExecute)
	http.Handle("/static/", http.FileServer(http.FS(staticDir)))

	// Port
	port := constants.AI_BUDDY_SERVER_PORT
	if s := os.Getenv(constants.AI_BUDDY_SERVER_PORT_ENV); s != "" {
		var err error
		port, err = strconv.Atoi(s)
		if err != nil {
			util.PrintError(fmt.Sprintf("Invalid port: %q", s))
			os.Exit(1)
		}
	}

	// Start server
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("ai-buddy server running on http://127.0.0.1%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		util.PrintError(err)
		os.Exit(1)
	}
}

// Handle pages

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	printLog(fmt.Sprintf("GET %s accessed", path))

	if path == "/" {
		path = "static/home.html"
	}

	loadPage(w, "", path)
}

func handleSettingsPage(w http.ResponseWriter, r *http.Request) {
	path := "static/settings.html"
	loadPage(w, "Settings", path)
}

func loadPage(w http.ResponseWriter, title string, path string) {
	pageData, err := staticDir.ReadFile(path)
	if err != nil {
			http.Error(w, fmt.Sprintf("Page not found: %s", path), http.StatusNotFound)
			return
	}
	page := Page{Title: title, Body: string(pageData)}

	t, err := template.ParseFS(staticDir, "static/template.html")
	if err != nil {
		util.PrintError(err)
		http.Error(w, "Template not found.", http.StatusNotFound)
		return
}
	w.WriteHeader(http.StatusOK)
  t.Execute(w, page)
}

// Hangle API requests

func handleGetProviders(w http.ResponseWriter, r *http.Request) {
	printLog("GET /providers accessed")

	providersHtml := `<option value="">Select a provider</option>`

	// Detect active providers
	if models.IsGooglePresent() {
		providersHtml += fmt.Sprintf(`<option value="%d">%s</option>`, models.MODEL_PROVIDER_GOOGLE, models.MODEL_PROVIDERS_NAMES[models.MODEL_PROVIDER_GOOGLE])
	}
	if models.IsOpenAIPresent() {
		providersHtml += fmt.Sprintf(`<option value="%d">%s</option>`, models.MODEL_PROVIDER_OPENAI, models.MODEL_PROVIDERS_NAMES[models.MODEL_PROVIDER_OPENAI])
	}
	if models.IsOllamaPresent() {
		providersHtml += fmt.Sprintf(`<option value="%d">%s</option>`, models.MODEL_PROVIDER_OLLAMA, models.MODEL_PROVIDERS_NAMES[models.MODEL_PROVIDER_OLLAMA])
	}

	fmt.Fprint(w, providersHtml)
}

func handleGetModels(w http.ResponseWriter, r *http.Request) {
	printLog("GET /models accessed")

	providerStr := r.URL.Query().Get("provider")
	provider, err := strconv.Atoi(providerStr)
	if err != nil {
		http.Error(w, "Provider not found", http.StatusNotFound)
		return
	}

	var modelsList []string

	if (provider == models.MODEL_PROVIDER_GOOGLE) {
		modelsList = models.MODEL_NAMES_GOOGLE
	} else if (provider == models.MODEL_PROVIDER_OPENAI) {
		modelsList = models.MODEL_NAMES_OPENAI
	} else if (provider == models.MODEL_PROVIDER_OLLAMA) {
		ollamaModels, err := models.GetOllamaModels()
		if err == nil {
			modelsList = ollamaModels
		}
	}

	modelsHtml := `<option value="">Select a model</option>`
	for _, m := range modelsList {
		modelsHtml += fmt.Sprintf(`<option value="%s">%s</option>`, strings.ReplaceAll(strings.ToLower(m), " ", "_"), m)
	}
	fmt.Fprint(w, modelsHtml)
}

func handleGetPatterns(w http.ResponseWriter, r *http.Request) {
	printLog("GET /patterns accessed")
	patternsHtml := `<option value="">Select a pattern</option>`
	patternsHtml += `<option value="no_pattern">NO PATTERN</option>`
	patterns, err := pattern.GetPatternList()
	if err == nil {
		for _, p := range patterns {
			patternsHtml += fmt.Sprintf(`<option value="%s">%s</option>`, p, p)
		}
		fmt.Fprint(w, patternsHtml)
	} else {
		http.Error(w, "Could not find pattern list", http.StatusNotFound)
		util.PrintError(err)
	}
}

func handleSettingsValue(w http.ResponseWriter, r *http.Request) {
	printLog("GET /settings_values accessed")

	settings := models.GetSettings()
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(settings)
}

func handleSaveSettings(w http.ResponseWriter, r *http.Request) {
	printLog("POST /save_settings accessed")

	r.ParseForm()

	googleApiKey := r.Form.Get("google_api_key")
	openaiApiKey := r.Form.Get("openai_api_key")

	settings := map[string]string{
		"googleApiKey": googleApiKey,
		"openaiApiKey": openaiApiKey,
	}

	res := models.SaveSettings(settings)
	if res {
		fmt.Fprint(w, "Settings saved.")
	} else {
		fmt.Fprint(w, "Could not save settings.")
	}
}

func handleGetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, constants.APP_VERSION)
}

func handleExecute(w http.ResponseWriter, r *http.Request) {
	printLog("POST /execute accessed")

	r.ParseForm()

	// Provider
	providerStr := r.Form.Get("provider")
	provider, err := strconv.Atoi(providerStr)
	if err != nil {
		http.Error(w, "Provider is required", http.StatusBadRequest)
		return
	}

	// Model
	modelName := r.Form.Get("model")
	if modelName == "" {
		http.Error(w, "Model is required", http.StatusBadRequest)
		return
	}

	// Pattern
	patternName := r.Form.Get("pattern")
	if patternName == "" {
		http.Error(w, "Pattern is required", http.StatusBadRequest)
		return
	}
	patternPrompt := ""
	if patternName == "no_pattern" {
		patternPrompt = ""
	} else {
		patternPrompt = pattern.GetPatternPrompt(patternName)
		if patternPrompt == "" {
			http.Error(w, "Pattern '"+ patternName + "' not found.", http.StatusNotFound)
			return
		}
	}

	// Prompt
	prompt := r.Form.Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt is required", http.StatusBadRequest)
		return
	}

	model := models.Model{Provider: provider, Name: modelName}
	model.SendPromptToModel(patternPrompt + prompt, w)
}

// Util

func printLog(text string) {
	if DEBUG {
		log.Print(text)
	}
}