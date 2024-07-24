package models

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cetorres/ai-buddy/util"
	"github.com/ollama/ollama/api"
	"github.com/ollama/ollama/envconfig"
)

func CreateOllamaGenerateStream(modelName string, prompt string) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		util.PrintError(err)
		os.Exit(1)
	}

	req := &api.GenerateRequest{
		Model:  modelName,
		Prompt: prompt,
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			util.PrintError("Could not find the Ollama server running on " + envconfig.Host.String() + ".")
		} else {
			util.PrintError(err)
		}
		os.Exit(1)
	}

	fmt.Println()
}