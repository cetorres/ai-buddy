package models

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cetorres/ai-buddy/util"
	"github.com/ollama/ollama/api"
	"github.com/ollama/ollama/envconfig"
)

func CreateOllamaGenerateStream(modelName string, prompt string, w http.ResponseWriter) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		if w != nil {
			fmt.Fprint(w, err)
		} else {
			util.PrintError(err)
			os.Exit(1)
		}
	}

	req := &api.GenerateRequest{
		Model:  modelName,
		Prompt: prompt,
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		if w != nil {
			fmt.Fprint(w, resp.Response)
		} else {
			fmt.Print(resp.Response)
		}
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			if w != nil {
				fmt.Fprint(w, "Could not find the Ollama server running on " + envconfig.Host.String() + ".")
			} else {
				util.PrintError("Could not find the Ollama server running on " + envconfig.Host.String() + ".")
			}
		} else {
			if w != nil {
				fmt.Fprint(w, err)
			} else {
				util.PrintError(err)
			}
		}
		if w == nil {
			os.Exit(1)
		}
	}
	if w == nil {
		fmt.Println()
	}
}