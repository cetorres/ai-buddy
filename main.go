package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const API_KEY_NAME = "GEMINI_API_KEY"
const TITLE = "AI Buddy 1.0 - Copyright © 2024 Carlos E. Torres (https://github.com/cetorres)"
var DESCRIPTION = fmt.Sprintf(`%s
An AI tool to help solving problems using a set of crowdsourced AI prompts.

Example usage:
	echo "Text to summarize..." | ai-buddy -p summarize
	ai-buddy -p summarize "Text to summarize..."
	cat MyEssayText.txt | ai-buddy -p analyze_claims
	pbpaste | ai-buddy -p summarize

Commands:
	-p, --pattern pattern_name prompt  Specify a pattern and send prompt to model. Requires pattern name and prompt (also receive via pipe).
	-l, --list                         List all available patterns.
	-v, --view pattern_name            View pattern prompt. Requires pattern name.
	-h, --help                         Show this help.

Uses the Google Gemini API:
	- Get your API key at https://aistudio.google.com/app/apikey
	- Set an environment variable: export %s=<your_key_here>

Patterns directory:
	- You can use the patterns directory in the same location of the binary (./patterns), this is by default.
	- Or you can set an environment variable if you want to move the binary to another directory.
	- Set the environment variable: export %s=<your_dir>/patterns`, TITLE, API_KEY_NAME, PATTERNS_DIR_ENV)

func main() {
	//
	// Check arguments
	//

	// Check API key, and show help
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

	// Try read text from forth argument
	if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && os.Args[3] != "" {
		executePatternCommand(os.Args[2], os.Args[3])
		os.Exit(0)
	} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
		printHelp()
		os.Exit(1)
	}

	// Try to read input from pipe
	if isInputFromPipe() {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		pipeString := string(stdin)
		if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && pipeString != "" {
			executePatternCommand(os.Args[2], pipeString)
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

func executePatternCommand(pattern string, text string) {
	patternPrompt := getPatternPrompt(pattern)
	if patternPrompt == "" {
		printError("Pattern '"+ pattern + "' not found.")
		os.Exit(1)
	}
	println(TITLE)
	println()
	sendPromptToModel(os.Getenv(API_KEY_NAME), patternPrompt + text)
}