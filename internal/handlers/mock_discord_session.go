package handlers

import (
	"github.com/bwmarrin/discordgo"
)

type MockDiscordSession struct {
	SentMessages []string
}

var _ DiscordSession = (*MockDiscordSession)(nil)

func (m *MockDiscordSession) ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error) {
	m.SentMessages = append(m.SentMessages, content)
	return &discordgo.Message{Content: content}, nil
}
