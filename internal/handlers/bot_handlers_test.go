package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
	"github.com/thornhall/chatgpt-discord-go/internal/constants"
	"github.com/thornhall/chatgpt-discord-go/internal/util"
)

func TestBotMessageHandler(t *testing.T) {
	tests := []struct {
		name          string
		expectedReply string
	}{
		{"WhompToThwomp", "Funny retort"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSession := &MockDiscordSession{}
			mockChat := &chatgptclient.MockChatService{
				MockResponse: tt.expectedReply,
			}

			var content []string
			content = append(content, "Test message")
			botMessage := util.BotMessage{
				ToBot:     constants.ThwompBot,
				FromBot:   constants.WhompBot,
				Content:   content,
				ChannelId: "FakeID",
			}
			botManager := &util.BotManager{}
			BotMessageHandler(mockSession, mockChat, botMessage, botManager)
			assert.Contains(t, mockSession.SentMessages, tt.expectedReply)
		})
	}
}
