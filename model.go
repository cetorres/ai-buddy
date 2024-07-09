package main

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func sendPromptToModel(apiKey string, prompt string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		printError(err)
		os.Exit(1)
	}
	
	defer client.Close()
	
	model := client.GenerativeModel("gemini-1.5-pro")
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
		printResponse(res)
	}
}
