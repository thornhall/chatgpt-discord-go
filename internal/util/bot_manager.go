package util

import (
	"github.com/bwmarrin/discordgo"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
)

type BotInstance struct {
	Session     *discordgo.Session
	ChatService chatgptclient.ChatService
}

type BotManager struct {
	Bots map[string]BotInstance
	Bus  chan BotMessage
}

type BotMessage struct {
	FromBot   string
	ToBot     string
	Content   []string
	ChannelId string
}
