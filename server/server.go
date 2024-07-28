package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/models"
	"github.com/cetorres/ai-buddy/pattern"
	"github.com/cetorres/ai-buddy/util"
)

//go:embed static/*
var	staticDir embed.FS

func CreateHTTPServer() {
	// Routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/providers", handleGetProviders)
	http.HandleFunc("/models", handleGetModels)
	http.HandleFunc("/patterns", handleGetPatterns)
	http.HandleFunc("POST /execute", handleExecute)
	http.Handle("/static/", http.FileServer(http.FS(staticDir)))

	// Port
	port := constants.AI_BUDDY_SERVER_PORT
	if s := os.Getenv(constants.AI_BUDDY_SERVER_PORT_ENV); s != "" {
		var err error
		port, err = strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Error: invalid port: %q - %s", s, err)
		}
	}

	// Start server
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("ai-buddy server running on http://127.0.0.1%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error %s", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("GET %s accessed", path)

	if path == "/" {
		path = "static/index.html"
	}

	data, err := staticDir.ReadFile(path)
	if err != nil {
			http.Error(w, fmt.Sprintf("Page not found: %s", path), http.StatusNotFound)
			return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleGetProviders(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /providers accessed")
	providers := models.MODEL_PROVIDERS
	providersHtml := `<option value="">Select a provider</option>`
	for i, p := range providers {
		providersHtml += fmt.Sprintf(`<option value="%d">%s</option>`, i+1, p)
	}
	fmt.Fprint(w, providersHtml)
}

func handleGetModels(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /models accessed")

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
		modelsList = []string{"llama3.1"}
	}

	modelsHtml := `<option value="">Select a model</option>`
	for _, m := range modelsList {
		modelsHtml += fmt.Sprintf(`<option value="%s">%s</option>`, strings.ReplaceAll(strings.ToLower(m), " ", "_"), m)
	}
	fmt.Fprint(w, modelsHtml)
}

func handleGetPatterns(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /patterns accessed")
	patternsHtml := `<option value="">Select a pattern</option>`
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

func handleExecute(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST /execute accessed")

	// provider=ollama&model=llama3.1&pattern=create_ideas&prompt=Meu%20prompt%20aqui

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
	patternPrompt := pattern.GetPatternPrompt(patternName)
	if patternPrompt == "" {
		http.Error(w, "Pattern '"+ patternName + "' not found.", http.StatusNotFound)
		return
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