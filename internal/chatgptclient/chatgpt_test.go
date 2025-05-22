package chatgptclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockChatService_GetChatGPTResponse(t *testing.T) {
	tests := []struct {
		name         string
		mockResponse string
		mockError    error
		prompt       string
		wantContent  string
		wantErr      bool
	}{
		{
			name:         "basic success",
			mockResponse: "Stop right there, criminal scum!",
			mockError:    nil,
			prompt:       "What's your duty?",
			wantContent:  "Stop right there, criminal scum!",
			wantErr:      false,
		},
		{
			name:         "error case",
			mockResponse: "",
			mockError:    assert.AnError,
			prompt:       "Trigger error",
			wantContent:  "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockChatService{
				MockResponse: tt.mockResponse,
				MockError:    tt.mockError,
			}

			resp, err := mock.GetChatGPTResponse(tt.prompt, "Test")

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.wantContent, resp.Choices[0].Message.Content)
			}
		})
	}
}
