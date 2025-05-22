package chatgptclient

import (
	openai "github.com/sashabaranov/go-openai"
)

type MockChatService struct {
	MockResponse string
	MockError    error
}

var _ ChatService = (*MockChatService)(nil)

func (m *MockChatService) GetChatGPTResponse(prompt string) (*openai.ChatCompletionResponse, error) {
	return &openai.ChatCompletionResponse{
		Choices: []openai.ChatCompletionChoice{
			{
				Message: openai.ChatCompletionMessage{
					Content: m.MockResponse,
				},
			},
		},
	}, m.MockError
}
