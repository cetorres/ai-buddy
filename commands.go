package main

import (
	"os"
	"strings"
)

func HelpCommand() {
	println(DESCRIPTION)
	if os.Getenv(GOOGLE_API_KEY_NAME) == "" {
		println()
		PrintError(GOOGLE_API_KEY_NAME + " is missing.")
	}
	if os.Getenv(OPENAI_API_KEY_NAME) == "" {
		println()
		PrintError(OPENAI_API_KEY_NAME + " is missing.")
	}
	os.Exit(0)
}

func ListCommand() {
	patterns, err := getPatternList()
	if patterns != nil {
		println("List of available patterns:\n")
		println(strings.Join(patterns, "\n"))
		os.Exit(0)
	} else {
		PrintError(err)
		os.Exit(1)
	}
}

func ListModelsCommand() {
	println("List of available models:\n")
	println("Google Gemini models:")
	println(strings.Join(MODEL_NAMES_GOOGLE, "\n"))
	println("\nOpenAI ChatGPT models:")
	println(strings.Join(MODEL_NAMES_OPENAI, "\n"))
	os.Exit(0)
}

func ViewCommand(pattern string) {
	patternPrompt := getPatternPrompt(pattern)
	if patternPrompt != "" {
		println("Pattern: " + pattern + "\n")
		println(patternPrompt)
		os.Exit(0)
	} else {
		PrintError("Pattern '"+ pattern + "' not found.")
		os.Exit(1)
	}
}

func PatternCommand(modelName string, pattern string, text string) {
	patternPrompt := getPatternPrompt(pattern)
	if patternPrompt == "" {
		PrintError("Pattern '"+ pattern + "' not found.")
		os.Exit(1)
	}

	if modelName != "" && !modelNameExists(modelName) {
		PrintError("Model '"+ modelName + "' not found.")
		os.Exit(1)
	}

	provider := MODEL_PROVIDER_GOOGLE
	if strings.Contains(modelName, "gpt") {
		provider = MODEL_PROVIDER_OPENAI
	}

	if modelName == "" {
		modelName = getDefaultModel()
	}

	model := Model{provider, modelName}
	model.sendPromptToModel(patternPrompt + text)
}