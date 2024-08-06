package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cetorres/ai-buddy/config"
	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/models"
	"github.com/cetorres/ai-buddy/pattern"
	"github.com/cetorres/ai-buddy/server"
	"github.com/cetorres/ai-buddy/util"
)

func HelpCommand() {
	print(constants.DESCRIPTION)
	println()
	conf := config.GetConfig()
	if conf.GoogleAPIKey == "" || conf.OpenAIAPIKey == "" {
		println()
	}
	if conf.GoogleAPIKey == "" {
		util.PrintWarning("Google Gemini API key is missing.")
	}
	if conf.OpenAIAPIKey == "" {
		util.PrintWarning("OpenAI API key is missing.")
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
		util.PrintError(fmt.Sprintf("ListCommand: %s", err))
		os.Exit(1)
	}
}

func ListModelsCommand() {
	println("List of available models:\n")
	
	println("Google Gemini models:")
	println(strings.Join(models.MODEL_NAMES_GOOGLE, "\n"))
	
	println("\nOpenAI ChatGPT models:")
	println(strings.Join(models.MODEL_NAMES_OPENAI, "\n"))

	println("\nAnthropic Claude models:")
	println(strings.Join(models.MODEL_NAMES_CLAUDE, "\n"))

	ollamaModels, err := models.GetOllamaModels()
	if err == nil {
		println("\nOllama models:")
		println(strings.Join(ollamaModels, "\n"))
	}

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
	conf := config.GetConfig()

	patternPrompt := pattern.GetPatternPrompt(patternName)
	if patternPrompt == "" {
		util.PrintError("Pattern '"+ patternName + "' not found.")
		os.Exit(1)
	}

	if provider == models.MODEL_PROVIDER_UNKNOWN {
		if conf.GoogleAPIKey != "" {
			provider = models.MODEL_PROVIDER_GOOGLE
		} else if conf.OpenAIAPIKey != "" {
			provider = models.MODEL_PROVIDER_OPENAI
		} else if conf.ClaudeAPIKey != "" {
			provider = models.MODEL_PROVIDER_CLAUDE
		}
	}
	
	if strings.Contains(modelName, "gpt") {
		provider = models.MODEL_PROVIDER_OPENAI
	} else if strings.Contains(modelName, "claude") {
		provider = models.MODEL_PROVIDER_CLAUDE
	}

	if provider == models.MODEL_PROVIDER_GOOGLE || provider == models.MODEL_PROVIDER_OPENAI || provider == models.MODEL_PROVIDER_CLAUDE {
		if provider == models.MODEL_PROVIDER_GOOGLE && conf.GoogleAPIKey == "" {
			util.PrintError("Google Gemini API key is missing.")
			os.Exit(1)
		}

		if provider == models.MODEL_PROVIDER_OPENAI && conf.OpenAIAPIKey == "" {
			util.PrintError("OpenAI API key is missing.")
			os.Exit(1)
		}

		if provider == models.MODEL_PROVIDER_CLAUDE && conf.ClaudeAPIKey == "" {
			util.PrintError("Anthropic Claude API key is missing.")
			os.Exit(1)
		}

		if modelName != "" && !models.ModelNameExists(modelName) {
			util.PrintError("Model '"+ modelName + "' not found.")
			os.Exit(1)
		}
	}

	if modelName == "" {
		modelName = models.GetDefaultModel(provider)
	}

	model := models.NewModel(provider, modelName)
	model.SendPromptToModel(patternPrompt + text, nil)
}

func ServeCommand(port int) {
	server.NewServer(port).Serve()
}

func SetupCommand() {
	conf := config.GetConfig()
	var googleApiKey, openaiApiKey, claudeApiKey string
	
	println("Welcome to ai-buddy. Let's get started.")
	println("\nCurrent configuration:")

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("-> Google Gemini API key: %s\n", conf.GoogleAPIKey)
	fmt.Printf("-> OpenAI API key: %s\n", conf.OpenAIAPIKey)
	fmt.Printf("-> Anthropic Claude API key: %s\n", conf.ClaudeAPIKey)

	println("\nTip: Leave field blank to not change it.")

	print("\n- Enter your Google Gemini API key: ")
  googleApiKey, err := reader.ReadString('\n')
	if err != nil {
		util.PrintError(err)
	}
	googleApiKey = strings.Trim(strings.TrimSpace(googleApiKey), "\n")
	if googleApiKey != "" {
		conf.GoogleAPIKey = googleApiKey
	}
	
	print("- Enter your OpenAI API key: ")
	openaiApiKey, err = reader.ReadString('\n')
	if err != nil {
		util.PrintError(err)
	}
	openaiApiKey = strings.Trim(strings.TrimSpace(openaiApiKey), "\n")
	if openaiApiKey != "" {
		conf.OpenAIAPIKey = openaiApiKey
	}

	print("- Enter your Anthropic Claude API key: ")
	claudeApiKey, err = reader.ReadString('\n')
	if err != nil {
		util.PrintError(err)
	}
	claudeApiKey = strings.Trim(strings.TrimSpace(claudeApiKey), "\n")
	if claudeApiKey != "" {
		conf.ClaudeAPIKey = claudeApiKey
	}
	
	// Save config
	config.SetConfig(conf)

	println()

	util.PrintSuccess("-> Configuration saved successfully.")

	// Copy patterns folder to config
	err = pattern.CopyPatternsDirToConfigDir()
	if err != nil {
		util.PrintError(fmt.Sprintf("Could not copy patterns directory: %s", err))
	} else {
		util.PrintSuccess("-> Patterns copied to config directory successfully.")
	}
}