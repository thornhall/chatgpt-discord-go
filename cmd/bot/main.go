package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/thornhall/chatgpt-discord-go/internal/chatgptclient"
	"github.com/thornhall/chatgpt-discord-go/internal/constants"
	"github.com/thornhall/chatgpt-discord-go/internal/handlers"
	"github.com/thornhall/chatgpt-discord-go/internal/util"
)

var envToBotRole = map[string]string{
	constants.OblivionGuardEnvVar: constants.OblivionGuardBot,
	constants.ThwompBotEnvVar:     constants.ThwompBot,
	constants.WhompBotEnvVar:      constants.WhompBot,
}

func startBot(botName string, token string, openAiToken string, manager *util.BotManager) (*discordgo.Session, error) {
	chatService := chatgptclient.NewChatService(openAiToken)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.MessageHandler(s, m, chatService, botName, manager)
	})

	go func() {
		for msg := range manager.Bus {
			if msg.ToBot == botName {
				handlers.BotMessageHandler(dg, chatService, msg, manager)
			}
		}
	}()

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	log.Printf("Bot %s is running", botName)
	return dg, nil
}

func main() {
	openAiKey := os.Getenv("OPENAI_API_KEY")
	if openAiKey == "" {
		_ = godotenv.Load(".env")
		openAiKey = os.Getenv("OPENAI_API_KEY")
	}
	if openAiKey == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}

	manager := &util.BotManager{
		Bots: make(map[string]*discordgo.Session),
		Bus:  make(chan util.BotMessage),
	}

	var sessions []*discordgo.Session

	for envVar, botName := range envToBotRole {
		token := os.Getenv(envVar)
		if token == "" {
			log.Fatalf("Missing discord API token for bot %s)", botName)
		}
		dg, err := startBot(botName, token, openAiKey, manager)
		if err != nil {
			log.Fatalf("Failed to start bot %s: %v", botName, err)
		}
		sessions = append(sessions, dg)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Received shutdown signal. Closing bots...")

	var wg sync.WaitGroup
	for _, session := range sessions {
		wg.Add(1)
		go func(s *discordgo.Session) {
			defer wg.Done()
			if err := s.Close(); err != nil {
				log.Printf("Error closing session: %v", err)
			}
		}(session)
	}
	wg.Wait()

	log.Println("All bots shut down gracefully. Exiting.")
}
