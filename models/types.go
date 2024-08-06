package models

import "github.com/sashabaranov/go-openai"

const (
	MODEL_PROVIDER_UNKNOWN int = iota
	MODEL_PROVIDER_GOOGLE
	MODEL_PROVIDER_OPENAI
	MODEL_PROVIDER_OLLAMA
	MODEL_PROVIDER_CLAUDE
)

var (
	MODEL_PROVIDERS_NAMES = []string{"Unknown", "Google Gemini", "OpenAI ChatGPT", "Ollama", "Anthropic Claude"}
	MODEL_NAMES_GOOGLE    = []string{"gemini-1.5-pro", "gemini-1.5-flash", "gemini-1.0-pro"}
	MODEL_NAMES_OPENAI    = []string{openai.GPT3Dot5Turbo, openai.GPT4, openai.GPT4o, openai.GPT4Turbo}
	MODEL_NAMES_CLAUDE    = []string{"claude-3-5-sonnet-20240620", "claude-3-opus-20240229", "claude-3-sonnet-20240229", "claude-3-haiku-20240307"}
)

type Model struct {
	Provider int
	Name string
}