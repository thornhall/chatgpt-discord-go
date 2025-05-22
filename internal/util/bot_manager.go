package util

import (
	"github.com/bwmarrin/discordgo"
)

type BotManager struct {
	Bots map[string]*discordgo.Session
	Bus  chan BotMessage
}

type BotMessage struct {
	FromBot   string
	ToBot     string
	Content   []string
	ChannelId string
}
