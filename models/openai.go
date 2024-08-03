package models

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cetorres/ai-buddy/config"
	"github.com/cetorres/ai-buddy/util"
	"github.com/sashabaranov/go-openai"
)

func CreateOpenAIChatStream(modelName string, prompt string, w http.ResponseWriter) {
	conf := config.GetConfig()
	client := openai.NewClient(conf.OpenAIAPIKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: modelName,
		MaxTokens: 0,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
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
	defer stream.Close()

	for {
		response, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			if w == nil {
				println()
			}
			return
		}

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

		if w != nil {
			w.Write([]byte(response.Choices[0].Delta.Content))
			w.(http.Flusher).Flush()
		} else {
			fmt.Printf(response.Choices[0].Delta.Content)
		}
		
	}
}

func IsOpenAIPresent() bool {
	conf := config.GetConfig()
	return conf.OpenAIAPIKey != ""
}