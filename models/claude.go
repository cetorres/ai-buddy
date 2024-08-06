package models

import (
	"fmt"
	"net/http"

	"github.com/cetorres/ai-buddy/config"
	"github.com/cetorres/ai-buddy/models/claude"
	"github.com/cetorres/ai-buddy/util"
)

func CreateClaudeChatCompletion(modelName string, prompt string, w http.ResponseWriter) {
	conf := config.GetConfig()
	client := claude.NewClient(conf.ClaudeAPIKey)

	req := &claude.ChatCompletionRequest{
    Model: modelName,
    Messages: []claude.Message{
      {Role: "user", Content: prompt},
    },
	}

	resp, err := client.ChatCompletion(req)
	if err != nil {
		if w != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			util.PrintError(err)
		} else {
			util.PrintError(err)
		}
	} else {
		if w != nil {
			fmt.Println(w, resp.Completion)
		} else {
			fmt.Println(resp.Completion)
		}
	}
}

func IsClaudePresent() bool {
	conf := config.GetConfig()
	return conf.ClaudeAPIKey != ""
}