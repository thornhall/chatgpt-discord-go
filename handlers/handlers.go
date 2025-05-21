package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/thornhall/chatgpt-discord-go/chatgptclient"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate, chatService *chatgptclient.ChatService) {
	if m.Author.Bot {
		return
	}

	content := strings.TrimSpace(m.Content)

	switch {
	case content == "!ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")

	case content == "!hello":
		s.ChannelMessageSend(m.ChannelID, "Hello there, "+m.Author.Username+"!")

	case strings.HasPrefix(content, "!echo "):
		text := strings.TrimPrefix(content, "!echo ")
		s.ChannelMessageSend(m.ChannelID, text)
	case content == "!tresspass":
		resp, err := chatService.GetChatGPTResponse("Respond as though a player has just committed a crime in the video game Elder Scrolls 4: Oblivion.")
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)
	case strings.HasPrefix(content, "!dialogue"):
		text := strings.TrimPrefix(content, "!dialogue")
		resp, err := chatService.GetChatGPTResponse(
			`Respond as though you're giving dialogue like an Oblivion NPC from the video game Elder Scrolls 4: 
		Oblivion but adapt it to the user's input, lean towards humor. Keep the responses under one paragraph. 
		Do not include quotation marks around your response. User input starts now: ` + text + ``)
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Message.Content)
	default:
	}
}
