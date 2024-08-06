package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIKey string
	BaseURL string
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		BaseURL: "https://api.anthropic.com/v1",
	}
}

type ChatCompletionRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens,omitempty"`
}

type Message struct {
	Role    string `json:"role"` // "user" or "assistant"
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	Completion string `json:"completion"`
}

func (c *Client) ChatCompletion(req *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	url := c.BaseURL + "/complete"

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshaling request body: %w", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	httpReq.Header.Set("content-type", "application/json")
	httpReq.Header.Set("anthropic-version", "2023-06-01")
	httpReq.Header.Set("x-api-key", c.APIKey)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var completionResp ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&completionResp); err != nil {
		return nil, fmt.Errorf("decoding response body: %w", err)
	}

	return &completionResp, nil
}