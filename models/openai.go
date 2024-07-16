package models

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/util"
	"github.com/sashabaranov/go-openai"
)

func CreateOpenAIChatStream(modelName string, prompt string) {
	apiKey := os.Getenv(constants.OPENAI_API_KEY_NAME)
	client := openai.NewClient(apiKey)
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
		util.PrintError(err)
		os.Exit(1)
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			println()
			return
		}

		if err != nil {
			util.PrintError(err)
			os.Exit(1)
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}