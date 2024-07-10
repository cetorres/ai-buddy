package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	MODEL_PROVIDER_GOOGLE = 0
	MODEL_PROVIDER_OPENAI = 1
)

var MODEL_NAMES_GOOGLE = []string{"gemini-1.5-pro", "gemini-1.5-flash", "gemini-1.0-pro"}
var MODEL_NAMES_OPENAI = []string{openai.GPT3Dot5Turbo, openai.GPT4, openai.GPT4o, openai.GPT4Turbo}

type Model struct {
	provider int
	name string
}

func (m Model) sendPromptToModel(prompt string) {
	if m.provider == MODEL_PROVIDER_GOOGLE {
		apiKey := os.Getenv(GOOGLE_API_KEY_NAME)
		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

		if err != nil {
			printError(err)
			os.Exit(1)
		}
		
		defer client.Close()
		
		model := client.GenerativeModel(m.name)
		session := model.StartChat()

		iter := session.SendMessageStream(ctx, genai.Text(prompt))
		for {
			res, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			printGoogleResponse(res)
		}

	} else if m.provider == MODEL_PROVIDER_OPENAI {
		apiKey := os.Getenv(OPENAI_API_KEY_NAME)
		c := openai.NewClient(apiKey)
		ctx := context.Background()

		req := openai.ChatCompletionRequest{
			Model: m.name,
			MaxTokens: 20,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Stream: true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		defer stream.Close()

		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				printError(err)
				os.Exit(1)
			}

			fmt.Printf(response.Choices[0].Delta.Content)
		}
	}
}

func getDefaultModel() string {
	if os.Getenv(DEFAULT_MODEL_ENV) != "" {
		return os.Getenv(DEFAULT_MODEL_ENV)
	}

	model := ""
	if os.Getenv(GOOGLE_API_KEY_NAME) != "" {
		model = MODEL_NAMES_GOOGLE[0]
	} else if os.Getenv(OPENAI_API_KEY_NAME) != "" {
		model = MODEL_NAMES_OPENAI[0]
	}
	return model 
}

func printGoogleResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Print(part)
			}
		}
	}
}
