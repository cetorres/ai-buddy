package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const API_KEY_NAME = "GEMINI_API_KEY"
const TITLE = "\nAI Buddy 1.0 - Copyright © 2024 Carlos E. Torres (https://github.com/cetorres)"
var DESCRIPTION = fmt.Sprintf(`%s
An AI tool to help solving problems using a set of crowdsourced AI prompts.

Example usage:
	$ echo "Text to summarize..." | ai-buddy -p summarize
	$ ai-buddy -p summarize "Text to summarize..."
	$ cat MyEssayText.txt | ai-buddy -p analyze_claims

Commands:
	-p or --pattern : Specify a pattern and send prompt to model. Requires pattern name and prompt.
	-l or --list    : List all available patterns.
	-v or --view    : View pattern prompt. Requires pattern name.
	-h or --help    : Show this help.

Uses the Google Gemini API:
	- Get your API key at https://aistudio.google.com/app/apikey
	- Set an environment variable: export %s=<your_key_here>`, TITLE, API_KEY_NAME)

func main() {
	// Check arguments, API key, and show help
	if (len(os.Args) < 2 || os.Getenv(API_KEY_NAME) == "") {
		printHelp()
		os.Exit(0)
	}

	// Check for help argument to show help
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printHelp()
		os.Exit(0)
	}

	// Check fo list argument
	if len(os.Args) == 2 && (os.Args[1] == "-l" || os.Args[1] == "--list") {
		patterns := getPatternList()
		println(TITLE)
		println("\nList of available patterns:\n")
		println(strings.Join(patterns, "\n"))
		os.Exit(0)
	}

	// Check for view argument
	if len(os.Args) == 3 && (os.Args[1] == "-v" || os.Args[1] == "--view") && os.Args[2] != "" {
		patternPrompt := getPatternPrompt(os.Args[2])
		println(TITLE)
		println("\nPattern: " + os.Args[2] + "\n")
		println(patternPrompt)
		os.Exit(0)
	}

	// Check for pattern argument

	// Try to read from pipe
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		printError(err)
	}
	pipeString := string(stdin)
	if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && pipeString != "" {
		executePatternCommand(os.Args[2], pipeString)
		os.Exit(0)
	}
	// Otherwise try read from forth argument
	if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && os.Args[3] != "" {
		executePatternCommand(os.Args[2], os.Args[3])
		os.Exit(0)
	}
}

func executePatternCommand(pattern string, text string) {
	patternPrompt := getPatternPrompt(pattern)
	if patternPrompt == "" {
		printError("Pattern does not exist.")
		os.Exit(1)
	}
	println(TITLE)
	println()
	sendPromptToModel(os.Getenv(API_KEY_NAME), patternPrompt + text)
}