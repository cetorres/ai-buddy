package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/util"
)

type OllamaError struct {
	Error string `json:"error"`
}

type OllamaResponse struct {
	Model string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response string `json:"response"`
	Done bool `json:"done"`
}

const (
	OLLAMA_LOCAL_URL = "http://localhost:11434"
	OLLAMA_GENERATE_API = "/api/generate"
)

func GetOllamaUrl() string {
	if os.Getenv(constants.OLLAMA_URL_ENV) != "" {
		return os.Getenv(constants.OLLAMA_URL_ENV)
	}
	return OLLAMA_LOCAL_URL 
}

func CreateOllamaGenerateStream(modelName string, prompt string) {
	values := map[string]string{"model": modelName, "prompt": prompt}
	json_data, err := json.Marshal(values)
	if err != nil {
		util.PrintError(err)
		os.Exit(1)
	}

	response, err := http.Post(GetOllamaUrl() + OLLAMA_GENERATE_API, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			util.PrintError("Could not find the Ollama server running on " + GetOllamaUrl() + ".")
		} else {
			util.PrintError(err)
		}
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		var ollamaError OllamaError
		err = json.NewDecoder(response.Body).Decode(&ollamaError)
		if err == nil {
			util.PrintError(ollamaError.Error)
		} else {
			util.PrintError("Ollama server response invalid: " + response.Status)
		}
		os.Exit(1)
	}

	for {
		var ollamaResponse OllamaResponse
		err = json.NewDecoder(response.Body).Decode(&ollamaResponse)
		if err != nil {
			util.PrintError(fmt.Sprintf("Could not decode JSON response: %s", err))
			os.Exit(1)
		}

		if err == io.EOF || ollamaResponse.Done {
			println()
			return
		}

		fmt.Printf(ollamaResponse.Response)
	}
}