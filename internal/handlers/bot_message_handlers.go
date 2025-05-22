package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
	"github.com/thornhall/chatgpt-discord-go/internal/constants"
	"github.com/thornhall/chatgpt-discord-go/internal/util"
)

type BotHandlerService interface {
	BotMessageHandler(s *discordgo.Session, chatService chatgptclient.ChatService, botMessage util.BotMessage, manager util.BotManager)
}

func BotMessageHandler(session DiscordSession, chatService chatgptclient.ChatService, botMessage util.BotMessage, manager *util.BotManager) {
	switch botMessage.ToBot {
	case constants.ThwompBot:
		ThwompBotMessageHandler(session, chatService, botMessage)
	case constants.WhompBot:
		WhompBotMessageHandler(session, chatService, botMessage)
	case constants.OblivionGuardBot:
		OblivionGuardBotMessageHandler(session, chatService, botMessage)
	}
}

func WhompBotMessageHandler(session DiscordSession, chatService chatgptclient.ChatService, botMessage util.BotMessage) {
}

func ThwompBotMessageHandler(session DiscordSession, chatService chatgptclient.ChatService, botMessage util.BotMessage) {
	rolePrompt := "You are Thwomp, a character from the video game series Mario Brothers by Nintendo. You have a rivalry with Whomp. You have received a message from Whomp, which will be in the user input. Respond in kind to his humor and/or insults."
	resp, err := chatService.GetChatGPTResponse(botMessage.Content[0], rolePrompt)
	if err != nil {
		log.Print("Request to chatGPT failed. " + err.Error() + "")
		return
	}
	session.ChannelMessageSend(botMessage.ChannelId, resp.Choices[0].Message.Content)
}

func OblivionGuardBotMessageHandler(session DiscordSession, chatService chatgptclient.ChatService, botMessage util.BotMessage) {

}
