package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
)

type DiscordSession interface {
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

type HandlerService interface {
	MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate, chatService *chatgptclient.ChatClient)
}

func MessageHandler(session DiscordSession, m *discordgo.MessageCreate, chatService chatgptclient.ChatService) {
	if m.Author.Bot {
		return
	}

	content := strings.TrimSpace(m.Content)

	switch {
	case content == "!ping":
		session.ChannelMessageSend(m.ChannelID, "Pong!")

	case content == "!hello":
		session.ChannelMessageSend(m.ChannelID, "Hello there, "+m.Author.Username+"!")

	case strings.HasPrefix(content, "!echo "):
		text := strings.TrimPrefix(content, "!echo ")
		session.ChannelMessageSend(m.ChannelID, text)
	case content == "!tresspass":
		resp, err := chatService.GetChatGPTResponse("Respond as though a player has just committed a crime in the video game Elder Scrolls 4: Oblivion.")
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		session.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)
	case strings.HasPrefix(content, "!dialogue"):
		text := strings.TrimPrefix(content, "!dialogue")
		resp, err := chatService.GetChatGPTResponse(
			`Respond as though you're giving dialogue like an Oblivion NPC from the video game Elder Scrolls 4: 
		Oblivion but adapt it to the user's input, lean towards humor. Keep the responses under one paragraph. 
		Do not include quotation marks around your response. If user input is empty just say something funny an oblivion guard would say. 
		User input starts now: ` + text + ``)
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		session.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)
	default:
	}
}
