package models

import "github.com/sashabaranov/go-openai"

const (
	MODEL_PROVIDER_GOOGLE = 1
	MODEL_PROVIDER_OPENAI = 2
	MODEL_PROVIDER_OLLAMA = 3
)

var MODEL_NAMES_GOOGLE = []string{"gemini-1.5-pro", "gemini-1.5-flash", "gemini-1.0-pro"}
var MODEL_NAMES_OPENAI = []string{openai.GPT3Dot5Turbo, openai.GPT4, openai.GPT4o, openai.GPT4Turbo}

type Model struct {
	Provider int
	Name string
}