package models

import (
	"net/http"
	"os"
	"slices"

	"github.com/cetorres/ai-buddy/constants"
)

func (m Model) SendPromptToModel(prompt string, w http.ResponseWriter) {
	// Google Gemini
	if m.Provider == MODEL_PROVIDER_GOOGLE {
		CreateGoogleMessageStream(m.Name, prompt, w)

	// OpenAI ChatGPT
	} else if m.Provider == MODEL_PROVIDER_OPENAI {
		CreateOpenAIChatStream(m.Name, prompt, w)

	// Ollama local server
	} else if m.Provider == MODEL_PROVIDER_OLLAMA {
		CreateOllamaGenerateStream(m.Name, prompt, w)
	}
}

func GetDefaultModel() string {
	if os.Getenv(constants.DEFAULT_MODEL_ENV) != "" {
		return os.Getenv(constants.DEFAULT_MODEL_ENV)
	}

	model := ""
	if os.Getenv(constants.GOOGLE_API_KEY_NAME) != "" {
		model = MODEL_NAMES_GOOGLE[0]
	} else if os.Getenv(constants.OPENAI_API_KEY_NAME) != "" {
		model = MODEL_NAMES_OPENAI[0]
	}
	return model 
}

func ModelNameExists(modelName string) bool {
	if !slices.Contains(MODEL_NAMES_GOOGLE, modelName) && !slices.Contains(MODEL_NAMES_OPENAI, modelName) {
		return false
	}
	return true
}
