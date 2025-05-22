package chatgptclient

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type ChatService interface {
	GetChatGPTResponse(promptText string) (*openai.ChatCompletionResponse, error)
}

type ChatClient struct {
	client *openai.Client
}

func NewChatService(apiKey string) *ChatClient {
	return &ChatClient{
		client: openai.NewClient(apiKey),
	}
}

func (c *ChatClient) GetChatGPTResponse(promptText string) (*openai.ChatCompletionResponse, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a guard from the Elder Scrolls video game Oblivion.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: promptText,
			},
		},
	}

	resp, err := c.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Printf("Chat completion error: %v", err)
		return nil, err
	}

	return &resp, nil
}
