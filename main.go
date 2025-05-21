package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/thornhall/chatgpt-discord-go/chatgptclient"
	"github.com/thornhall/chatgpt-discord-go/handlers"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	apiKey := os.Getenv("OPENAI_API_KEY")

	if token == "" {
		// If not already present in the environment, we're probably running it locally with a .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Unable to load .env file")
		}
		token = os.Getenv("DISCORD_BOT_TOKEN")
		apiKey = os.Getenv("OPENAI_API_KEY")
	}

	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN not set")
	}

	chatService := chatgptclient.NewChatService(apiKey)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.MessageHandler(s, m, chatService)
	})
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	log.Println("Bot is running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	dg.Close()
}
