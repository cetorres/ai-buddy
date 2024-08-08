package models

import (
	"net/http"
	"slices"

	"github.com/cetorres/ai-buddy/config"
)

func NewModel(provider int, modelName string) *Model {
	return &Model{
		Provider: provider,
		Name: modelName,
	}
}

func (m Model) SendPromptToModel(prompt string, w http.ResponseWriter) {
	// Google Gemini
	if m.Provider == MODEL_PROVIDER_GOOGLE {
		m.CreateGoogleMessageStream(prompt, w)

	// OpenAI ChatGPT
	} else if m.Provider == MODEL_PROVIDER_OPENAI {
		m.CreateOpenAIChatStream(prompt, w)

	// Anthropic Claude
	} else if m.Provider == MODEL_PROVIDER_CLAUDE {
		m.CreateClaudeMessageStream(prompt, w)

	// Ollama local server
	} else if m.Provider == MODEL_PROVIDER_OLLAMA {
		m.CreateOllamaGenerateStream(prompt, w)
	}
}

func GetDefaultModel(provider int) string {
	conf := config.GetConfig()

	model := ""
	if provider == MODEL_PROVIDER_GOOGLE && conf.GoogleAPIKey != "" {
		model = MODEL_NAMES_GOOGLE[0]
	} else if provider == MODEL_PROVIDER_OPENAI && conf.OpenAIAPIKey != "" {
		model = MODEL_NAMES_OPENAI[0]
	} else if provider == MODEL_PROVIDER_CLAUDE && conf.ClaudeAPIKey != "" {
		model = MODEL_NAMES_CLAUDE[0]
	}
	return model 
}

func ModelNameExists(modelName string) bool {
	if !slices.Contains(MODEL_NAMES_GOOGLE, modelName) && !slices.Contains(MODEL_NAMES_OPENAI, modelName) && !slices.Contains(MODEL_NAMES_CLAUDE, modelName) {
		return false
	}
	return true
}
