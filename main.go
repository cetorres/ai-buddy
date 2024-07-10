package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	//
	// Check arguments
	//

	// Check API key, and show help
	if (len(os.Args) < 2 || (os.Getenv(GOOGLE_API_KEY_NAME) == "" && os.Getenv(OPENAI_API_KEY_NAME) == "")) {
		printHelp()
		os.Exit(0)
	}

	// Check for help argument to show help
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printHelp()
		os.Exit(0)
	}

	// Check for list argument
	if len(os.Args) >= 2 && (os.Args[1] == "-l" || os.Args[1] == "--list") {
		patterns, err := getPatternList()
		if patterns != nil {
			println(TITLE)
			println("\nList of available patterns:\n")
			println(strings.Join(patterns, "\n"))
			os.Exit(0)
		} else {
			printError(err)
			os.Exit(1)
		}
	}

	// Check for list models argument
	if len(os.Args) >= 2 && (os.Args[1] == "-lm" || os.Args[1] == "--list-models") {
		println(TITLE)
		println("\nGoogle Gemini models:")
		println(strings.Join(MODEL_NAMES_GOOGLE, "\n"))
		println("\nOpenAI ChatGPT models:")
		println(strings.Join(MODEL_NAMES_OPENAI, "\n"))
		os.Exit(0)
	}

	// Check for view argument
	if len(os.Args) >= 3 && (os.Args[1] == "-v" || os.Args[1] == "--view") && os.Args[2] != "" {
		pattern := os.Args[2]
		patternPrompt := getPatternPrompt(pattern)
		if patternPrompt != "" {
			println(TITLE)
			println("\nPattern: " + pattern + "\n")
			println(patternPrompt)
			os.Exit(0)
		} else {
			printError("Pattern '"+ pattern + "' not found.")
			os.Exit(1)
		}
	} else if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--view") {
		printHelp()
		os.Exit(1)
	}

	//
	// Check for pattern argument
	//

	// Try to read input from pipe
	if isInputFromPipe() {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		pipeString := string(stdin)
		if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" && pipeString != "" {
			executePatternCommand(os.Args[4], os.Args[2], pipeString)
			os.Exit(0)
		} else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && pipeString != "" {
			printHelp()
			os.Exit(1)
		} else if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && pipeString != "" {
			executePatternCommand("", os.Args[2], pipeString)
			os.Exit(0)
		} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
			printHelp()
			os.Exit(1)
		}
	} else {
		// Try read text from argument
		if len(os.Args) == 6 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" && os.Args[5] != "" {
			executePatternCommand(os.Args[4], os.Args[2], os.Args[5])
			os.Exit(0)
		} else if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" {
			printHelp()
			os.Exit(1)
		} else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && os.Args[3] != "" {
			executePatternCommand("", os.Args[2], os.Args[3])
			os.Exit(0)
		} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
			printHelp()
			os.Exit(1)
		}
	}
	
	// Show help if cannot identify arguments
	printHelp()
	os.Exit(1)
}

func executePatternCommand(modelName string, pattern string, text string) {
	patternPrompt := getPatternPrompt(pattern)
	if patternPrompt == "" {
		printError("Pattern '"+ pattern + "' not found.")
		os.Exit(1)
	}

	provider := MODEL_PROVIDER_GOOGLE
	if strings.Contains(modelName, "gpt") {
		provider = MODEL_PROVIDER_OPENAI
	}

	if modelName == "" {
		modelName = getDefaultModel()
	}

	println(TITLE)
	println()

	model := Model{provider, modelName}
	model.sendPromptToModel(patternPrompt + text)
}