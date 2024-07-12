package main

import (
	"io"
	"os"
)

func main() {
	// Check API keys and number of arguments, and show help
	if (len(os.Args) < 2 || (os.Getenv(GOOGLE_API_KEY_NAME) == "" && os.Getenv(OPENAI_API_KEY_NAME) == "")) {
		HelpCommand()
		os.Exit(0)
	}

	// Check for help argument to show help
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		HelpCommand()
		os.Exit(0)
	}

	// Check for list argument
	if len(os.Args) >= 2 && (os.Args[1] == "-l" || os.Args[1] == "--list") {
		ListCommand()
	}

	// Check for list models argument
	if len(os.Args) >= 2 && (os.Args[1] == "-lm" || os.Args[1] == "--list-models") {
		ListModelsCommand()
	}

	// Check for view argument
	if len(os.Args) >= 3 && (os.Args[1] == "-v" || os.Args[1] == "--view") && os.Args[2] != "" {
		ViewCommand(os.Args[2])
	} else if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--view") {
		PrintError("A pattern was not specified.")
		os.Exit(1)
	}

	//
	// Check for pattern argument
	//

	// Try to read input from pipe
	if IsInputFromPipe() {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			PrintError(err)
			os.Exit(1)
		}
		pipeString := string(stdin)
		if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" && pipeString != "" {
			PatternCommand(os.Args[4], os.Args[2], pipeString)
			os.Exit(0)
		} else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && pipeString != "" {
			PrintError("A model was not specified.")
			os.Exit(1)
		} else if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && pipeString != "" {
			PatternCommand("", os.Args[2], pipeString)
			os.Exit(0)
		} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
			PrintError("A pattern was not specified.")
			os.Exit(1)
		}
	} else {
		// Try read text from argument
		if len(os.Args) == 6 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" && os.Args[5] != "" {
			PatternCommand(os.Args[4], os.Args[2], os.Args[5])
			os.Exit(0)
		} else if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" {
			PrintError("A prompt was not entered.")
			os.Exit(1)
		} else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && os.Args[3] != "" {
			PatternCommand("", os.Args[2], os.Args[3])
			os.Exit(0)
		} else if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" {
			PrintError("A prompt was not entered.")
			os.Exit(1)
		} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
			PrintError("A pattern was not specified.")
			os.Exit(1)
		}
	}
	
	// Show help if cannot identify arguments
	HelpCommand()
	os.Exit(1)
}
