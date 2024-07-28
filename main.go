package main

import (
	"io"
	"os"
	"slices"

	"github.com/cetorres/ai-buddy/commands"
	"github.com/cetorres/ai-buddy/models"
	"github.com/cetorres/ai-buddy/util"
)

func main() {
	// Check API keys and number of arguments, and show help
	if len(os.Args) < 2 {
		commands.HelpCommand()
		os.Exit(0)
	}

	// Check for help argument to show help
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		commands.HelpCommand()
		os.Exit(0)
	}

	// Check for webui command
	if len(os.Args) == 2 && (os.Args[1] == "-w" || os.Args[1] == "--webui") {
		commands.ServeCommand()
		os.Exit(0)
	}

	// Check for list argument
	if len(os.Args) >= 2 && (os.Args[1] == "-l" || os.Args[1] == "--list") {
		commands.ListCommand()
	}

	// Check for list models argument
	if len(os.Args) >= 2 && (os.Args[1] == "-lm" || os.Args[1] == "--list-models") {
		commands.ListModelsCommand()
	}

	// Check for view argument
	if len(os.Args) >= 3 && (os.Args[1] == "-v" || os.Args[1] == "--view") && os.Args[2] != "" {
		commands.ViewCommand(os.Args[2])
	} else if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--view") {
		util.PrintError("A pattern was not specified.")
		os.Exit(1)
	}

	//
	// Check for pattern argument
	//

	// Try to read input from pipe
	if util.IsInputFromPipe() {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			util.PrintError(err)
			os.Exit(1)
		}
		pipeString := string(stdin)
		if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" && pipeString != "" {
			commands.PatternCommand(os.Args[4], os.Args[2], pipeString, models.MODEL_PROVIDER_UNKNOWN)
			os.Exit(0)
		} else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && pipeString != "" {
			util.PrintError("A model was not specified.")
			os.Exit(1)
		} else if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && pipeString != "" {
			commands.PatternCommand("", os.Args[2], pipeString, models.MODEL_PROVIDER_UNKNOWN)
			os.Exit(0)
		} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
			util.PrintError("A pattern was not specified.")
			os.Exit(1)
		} else if slices.Contains(os.Args, "-o") || slices.Contains(os.Args, "-ollama") {
			if len(os.Args) == 6 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-o" || os.Args[3] == "--ollama") && (os.Args[4] == "-m" || os.Args[4] == "--model") && os.Args[5] != "" && pipeString != "" {
				commands.PatternCommand(os.Args[5], os.Args[2], pipeString, models.MODEL_PROVIDER_OLLAMA)
				os.Exit(0)
			} else if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-o" || os.Args[3] == "--ollama") && (os.Args[4] == "-m" || os.Args[4] == "--model") {
				util.PrintError("A model was not specified.")
				os.Exit(1)
			}	else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-o" || os.Args[3] == "--ollama") {
				util.PrintError("A model was not specified.")
				os.Exit(1)
			}
		}
	} else {
		// Try read text from argument
		if len(os.Args) == 6 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" && os.Args[5] != "" {
			commands.PatternCommand(os.Args[4], os.Args[2], os.Args[5], models.MODEL_PROVIDER_UNKNOWN)
			os.Exit(0)
		} else if len(os.Args) == 5 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-m" || os.Args[3] == "--model") && os.Args[4] != "" {
			util.PrintError("A prompt was not entered.")
			os.Exit(1)
		} else if len(os.Args) == 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && os.Args[3] != "" {
			commands.PatternCommand("", os.Args[2], os.Args[3], models.MODEL_PROVIDER_UNKNOWN)
			os.Exit(0)
		} else if len(os.Args) == 3 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" {
			util.PrintError("A prompt was not entered.")
			os.Exit(1)
		} else if len(os.Args) == 2 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") {
			util.PrintError("A pattern was not specified.")
			os.Exit(1)
		} else if slices.Contains(os.Args, "-o") || slices.Contains(os.Args, "-ollama") {
			if len(os.Args) == 7 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-o" || os.Args[3] == "--ollama") && (os.Args[4] == "-m" || os.Args[4] == "--model") && os.Args[5] != "" && os.Args[6] != "" {
				commands.PatternCommand(os.Args[5], os.Args[2], os.Args[6], models.MODEL_PROVIDER_OLLAMA)
				os.Exit(0)
			} else if len(os.Args) == 6 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-o" || os.Args[3] == "--ollama") && (os.Args[4] == "-m" || os.Args[4] == "--model") && os.Args[5] != "" {
				util.PrintError("A model was not specified.")
				os.Exit(1)
			} else if len(os.Args) >= 4 && (os.Args[1] == "-p" || os.Args[1] == "--pattern") && os.Args[2] != "" && (os.Args[3] == "-o" || os.Args[3] == "--ollama") {
				util.PrintError("A model was not specified.")
				os.Exit(1)
			}
		}
	}
	
	// Show help if cannot identify arguments
	commands.HelpCommand()
	os.Exit(1)
}
