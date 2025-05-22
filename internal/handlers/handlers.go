package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
	"github.com/thornhall/chatgpt-discord-go/internal/constants"
)

type DiscordSession interface {
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

type HandlerService interface {
	MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate, chatService *chatgptclient.ChatClient)
}

func MessageHandler(session DiscordSession, m *discordgo.MessageCreate, chatService chatgptclient.ChatService, botName string) {
	if m.Author.Bot {
		return
	}
	content := strings.TrimSpace(m.Content)
	switch botName {
	case constants.OblivionGuardBot:
		OblivionGuardMessageHandler(session, m, chatService, content)
	case constants.ThwompBot:
		ThwompMessageHandler(session, m, chatService, content)
	case constants.WhompBot:
		WhompMessageHandler(session, m, chatService, content)
	}
}

func WhompMessageHandler(session DiscordSession, message *discordgo.MessageCreate, chatService chatgptclient.ChatService, content string) {
	rolePrompt := "You are a Whomp, a character from the video game series Mario Brothers by Nintendo. You have a rivalry with another AI, Thwomp Bot, and you're always competing with him. Do not give responses with quotes around them, try to lean towards using humor, and keep responses a paragraph or shorter."
	switch {
	case content == "!ping":
		session.ChannelMessageSend(message.ChannelID, "Pong!")

	case content == "!hello":
		session.ChannelMessageSend(message.ChannelID, "Hello there, "+message.Author.Username+"!")

	case strings.HasPrefix(content, "!echo "):
		text := strings.TrimPrefix(content, "!echo ")
		session.ChannelMessageSend(message.ChannelID, text)
	case content == "!tresspass":
		resp, err := chatService.GetChatGPTResponse("Respond as though a player has just committed a crime in the video game Elder Scrolls 4: Oblivion.", rolePrompt)
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		session.ChannelMessageSend(message.ChannelID, resp.Choices[0].Message.Content)
	case strings.HasPrefix(content, "!dialogue"):
		return
	default:
	}
}

func ThwompMessageHandler(session DiscordSession, message *discordgo.MessageCreate, chatService chatgptclient.ChatService, content string) {
	rolePrompt := "You are a Thwomp, a character from the video game series Mario Brothers by Nintendo. You have a rivalry with another AI, Whomp Bot, and you're always competing with him. Do not give responses with quotes around them, try to lean towards using humor, and keep responses a paragraph or shorter."
	switch {
	case content == "!ping":
		session.ChannelMessageSend(message.ChannelID, "Pong!")

	case content == "!hello":
		session.ChannelMessageSend(message.ChannelID, "Hello there, "+message.Author.Username+"!")

	case strings.HasPrefix(content, "!echo "):
		text := strings.TrimPrefix(content, "!echo ")
		session.ChannelMessageSend(message.ChannelID, text)
	case content == "!tresspass":
		resp, err := chatService.GetChatGPTResponse("Respond as though a player has just committed a crime in the video game Elder Scrolls 4: Oblivion.", rolePrompt)
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		session.ChannelMessageSend(message.ChannelID, resp.Choices[0].Message.Content)
	case strings.HasPrefix(content, "!dialogue"):
		return
	default:
	}
}

func OblivionGuardMessageHandler(session DiscordSession, message *discordgo.MessageCreate, chatService chatgptclient.ChatService, content string) {
	rolePrompt :=
		`Respond as though you're giving dialogue like an Oblivion NPC from the video game Elder Scrolls 4: 
		Oblivion but adapt it to the user's input, lean towards humor. Keep the responses under one paragraph. 
		Do not include quotation marks around your response.`
	switch {
	case content == "!ping":
		session.ChannelMessageSend(message.ChannelID, "Pong!")

	case content == "!hello":
		session.ChannelMessageSend(message.ChannelID, "Hello there, "+message.Author.Username+"!")

	case strings.HasPrefix(content, "!echo "):
		text := strings.TrimPrefix(content, "!echo ")
		session.ChannelMessageSend(message.ChannelID, text)
	case content == "!tresspass":
		resp, err := chatService.GetChatGPTResponse("Respond as though a player has just committed a crime in the video game Elder Scrolls 4: Oblivion.", rolePrompt)
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		session.ChannelMessageSend(message.ChannelID, resp.Choices[0].Message.Content)
	case strings.HasPrefix(content, "!dialogue"):
		text := strings.TrimPrefix(content, "!dialogue")
		resp, err := chatService.GetChatGPTResponse(text, rolePrompt)
		if err != nil {
			log.Print("Request to chatGPT failed. " + err.Error() + "")
			return
		}
		session.ChannelMessageSend(message.ChannelID, resp.Choices[0].Message.Content)
	default:
	}
}
