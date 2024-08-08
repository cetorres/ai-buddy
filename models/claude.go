package models

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/cetorres/ai-buddy/config"
	"github.com/cetorres/ai-buddy/util"
	"github.com/liushuangls/go-anthropic/v2"
)

func (m Model) CreateClaudeMessageStream(prompt string, w http.ResponseWriter) {
	conf := config.GetConfig()

	client := anthropic.NewClient(conf.ClaudeAPIKey)
	
	resp, err := client.CreateMessagesStream(context.Background(),  anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model: m.Name,
			Messages: []anthropic.Message{
				anthropic.NewUserTextMessage(prompt),
			},
			MaxTokens: 0,
		},
		OnContentBlockDelta: func(data anthropic.MessagesEventContentBlockDeltaData) {
			if w != nil {
				w.Write([]byte(*data.Delta.Text))
				w.(http.Flusher).Flush()
			} else {
				fmt.Printf("%s", *data.Delta.Text)
			}
		},
	})

	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			if w != nil {
				http.Error(w, fmt.Sprintf("ERROR: %s. %s", e.Type, e.Message), http.StatusInternalServerError)
				util.PrintError(fmt.Sprintf("%s. %s", e.Type, e.Message))
			} else {
				util.PrintError(fmt.Sprintf("%s. %s", e.Type, e.Message))
				os.Exit(1)
			}
		} else {
			if w != nil {
				http.Error(w, fmt.Sprintf("ERROR: %v", err), http.StatusInternalServerError)
				util.PrintError(fmt.Sprintf("%v", err))
			} else {
				util.PrintError(fmt.Sprintf("%v", err))
				os.Exit(1)
			}
    }
		return
	}

	if w != nil {
		w.Write([]byte(*resp.Content[0].Text))
		w.(http.Flusher).Flush()
	} else {
		fmt.Printf(*resp.Content[0].Text)
	}


	// client := claude.NewClient(conf.ClaudeAPIKey)

	// req := &claude.ChatCompletionRequest{
  //   Model: m.Name,
  //   Messages: []claude.Message{
  //     {Role: "user", Content: prompt},
  //   },
	// }

	// resp, err := client.ChatCompletion(req)
	// if err != nil {
		// if w != nil {
		// 	http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		// 	util.PrintError(err)
		// } else {
		// 	util.PrintError(err)
		// }
	// } else {
		// if w != nil {
		// 	fmt.Println(w, resp.Completion)
		// } else {
		// 	fmt.Println(resp.Completion)
		// }
	// }
}

func IsClaudePresent() bool {
	conf := config.GetConfig()
	return conf.ClaudeAPIKey != ""
}