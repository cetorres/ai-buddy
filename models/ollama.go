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
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			util.PrintError(err)
			return
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
			w.Write([]byte(resp.Response))
			w.(http.Flusher).Flush()			
		} else {
			fmt.Print(resp.Response)
		}
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			if w != nil {
				http.Error(w, "Could not find the Ollama server running on " + envconfig.Host().String() + ".", http.StatusInternalServerError)
			}
			util.PrintError("Could not find the Ollama server running on " + envconfig.Host().String() + ".")
		} else {
			if w != nil {
				http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			}
			util.PrintError(err)
		}
		if w == nil {
			os.Exit(1)
		}
	}
	if w == nil {
		fmt.Println()
	}
}

func IsOllamaPresent() bool {
	_, err := GetOllamaModels()
	return err == nil
}

func GetOllamaModels() ([]string, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	list, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	models := []string{}
	for _, m := range list.Models {
		models = append(models, m.Name)
	}

	return models, nil
}