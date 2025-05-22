package handlers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
	"github.com/thornhall/chatgpt-discord-go/internal/constants"
	"github.com/thornhall/chatgpt-discord-go/internal/util"
)

func TestMessageHandler(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		expectedReply string
	}{
		{"ping", "!ping", "Pong!"},
		{"hello", "!hello", "Hello there, testuser!"},
		{"echo", "!echo Hi there", "Hi there"},
		{"tresspass", "!tresspass", "Stop right there, criminal scum!"},
		{"dialogue", "!dialogue You dare challenge me?", "Funny Oblivion line"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSession := &MockDiscordSession{}
			mockChat := &chatgptclient.MockChatService{
				MockResponse: tt.expectedReply,
			}

			message := &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Content:   tt.content,
					Author:    &discordgo.User{Username: "testuser", Bot: false},
					ChannelID: "channel123",
				},
			}

			botManager := &util.BotManager{}
			MessageHandler(mockSession, message, mockChat, constants.OblivionGuardBot, botManager)
			assert.Contains(t, mockSession.SentMessages, tt.expectedReply)
		})
	}
}
