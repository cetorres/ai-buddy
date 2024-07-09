package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
)

func printError(err any) {
	fmt.Println(TITLE)
	color.Red(fmt.Sprintf("ERROR: %s", err))
}

func printHelp() {
	fmt.Println(DESCRIPTION)
	if os.Getenv(API_KEY_NAME) == "" {
		color.Red("\nERROR: " + API_KEY_NAME + " is missing.")
		os.Exit(1)
	}
	os.Exit(0)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Print(part)
			}
		}
	}
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}