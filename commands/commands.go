package commands

import (
	"os"
	"strings"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/models"
	"github.com/cetorres/ai-buddy/pattern"
	"github.com/cetorres/ai-buddy/server"
	"github.com/cetorres/ai-buddy/util"
)

func HelpCommand() {
	println(constants.DESCRIPTION)
	println()
	if os.Getenv(constants.GOOGLE_API_KEY_NAME) == "" {
		util.PrintWarning(constants.GOOGLE_API_KEY_NAME + " is missing.")
	}
	if os.Getenv(constants.OPENAI_API_KEY_NAME) == "" {
		util.PrintWarning(constants.OPENAI_API_KEY_NAME + " is missing.")
	}
	os.Exit(0)
}

func ListCommand() {
	patterns, err := pattern.GetPatternList()
	if patterns != nil {
		println("List of available patterns:\n")
		println(strings.Join(patterns, "\n"))
		os.Exit(0)
	} else {
		util.PrintError(err)
		os.Exit(1)
	}
}

func ListModelsCommand() {
	println("List of available models:\n")
	println("Google Gemini models:")
	println(strings.Join(models.MODEL_NAMES_GOOGLE, "\n"))
	println("\nOpenAI ChatGPT models:")
	println(strings.Join(models.MODEL_NAMES_OPENAI, "\n"))
	os.Exit(0)
}

func ViewCommand(patternName string) {
	patternPrompt := pattern.GetPatternPrompt(patternName)
	if patternPrompt != "" {
		println("Pattern: " + patternName + "\n")
		println(patternPrompt)
		os.Exit(0)
	} else {
		util.PrintError("Pattern '"+ patternName + "' not found.")
		os.Exit(1)
	}
}

func PatternCommand(modelName string, patternName string, text string, provider int) {
	patternPrompt := pattern.GetPatternPrompt(patternName)
	if patternPrompt == "" {
		util.PrintError("Pattern '"+ patternName + "' not found.")
		os.Exit(1)
	}

	if provider == 0 {
		provider = models.MODEL_PROVIDER_GOOGLE
	}
	
	if strings.Contains(modelName, "gpt") {
		provider = models.MODEL_PROVIDER_OPENAI
	}

	if provider == models.MODEL_PROVIDER_GOOGLE || provider == models.MODEL_PROVIDER_OPENAI {
		if provider == models.MODEL_PROVIDER_GOOGLE && os.Getenv(constants.GOOGLE_API_KEY_NAME) == "" {
			util.PrintError(constants.GOOGLE_API_KEY_NAME + " is missing.")
			os.Exit(1)
		}

		if provider == models.MODEL_PROVIDER_OPENAI && os.Getenv(constants.OPENAI_API_KEY_NAME) == "" {
			util.PrintError(constants.OPENAI_API_KEY_NAME + " is missing.")
			os.Exit(1)
		}

		if modelName != "" && !models.ModelNameExists(modelName) {
			util.PrintError("Model '"+ modelName + "' not found.")
			os.Exit(1)
		}
	}

	if modelName == "" {
		modelName = models.GetDefaultModel()
	}

	model := models.Model{Provider: provider, Name: modelName}
	model.SendPromptToModel(patternPrompt + text, nil)
}

func ServeCommand() {
	server.CreateHTTPServer()
}