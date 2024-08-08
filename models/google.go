package models

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cetorres/ai-buddy/config"
	"github.com/cetorres/ai-buddy/util"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func (m Model) CreateGoogleMessageStream(prompt string, w http.ResponseWriter) {
	conf := config.GetConfig()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.GoogleAPIKey))

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
	
	model := client.GenerativeModel(m.Name)
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

		printGoogleResponse(response, w)
	}
}

func printGoogleResponse(resp *genai.GenerateContentResponse, w http.ResponseWriter) {
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
	conf := config.GetConfig()
	return conf.GoogleAPIKey != ""
}