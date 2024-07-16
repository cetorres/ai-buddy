package models

import (
	"context"
	"fmt"
	"os"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/util"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func CreateGoogleMessageStream(modelName string, prompt string) {
	apiKey := os.Getenv(constants.GOOGLE_API_KEY_NAME)
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		util.PrintError(err)
		os.Exit(1)
	}
	
	defer client.Close()
	
	model := client.GenerativeModel(modelName)
	session := model.StartChat()

	iter := session.SendMessageStream(ctx, genai.Text(prompt))
	for {
		response, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			util.PrintError(err)
			os.Exit(1)
		}

		PrintGoogleResponse(response)
	}
}

func PrintGoogleResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Print(part)
			}
		}
	}
}