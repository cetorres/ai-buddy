package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func printError(err any) {
	color.Red(fmt.Sprintf("ERROR: %s", err))
}

func printHelp() {
	fmt.Println(DESCRIPTION)
	if os.Getenv(GOOGLE_API_KEY_NAME) == "" {
		color.Red("\nERROR: " + GOOGLE_API_KEY_NAME + " is missing.")
		os.Exit(1)
	}
	if os.Getenv(OPENAI_API_KEY_NAME) == "" {
		color.Red("\nERROR: " + OPENAI_API_KEY_NAME + " is missing.")
		os.Exit(1)
	}
	os.Exit(0)
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}