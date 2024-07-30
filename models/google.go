package models

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cetorres/ai-buddy/constants"
	"github.com/cetorres/ai-buddy/util"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func CreateGoogleMessageStream(modelName string, prompt string, w http.ResponseWriter) {
	apiKey := os.Getenv(constants.GOOGLE_API_KEY_NAME)
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		if w != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		}
		util.PrintError(err)
		if w == nil {
			os.Exit(1)
		} else {
			return
		}
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
			if w != nil {
				http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
				util.PrintError(err)
				return
			} else {
				util.PrintError(err)			
				os.Exit(1)
			}
		}

		PrintGoogleResponse(response, w)
	}
}

func PrintGoogleResponse(resp *genai.GenerateContentResponse, w http.ResponseWriter) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if w != nil {
					w.Write([]byte(fmt.Sprintf("%s", part)))
					w.(http.Flusher).Flush()	
				} else {
					fmt.Print(part)
				}
			}
		}
	}
}

func IsGooglePresent() bool {
	return os.Getenv(constants.GOOGLE_API_KEY_NAME) != ""
}