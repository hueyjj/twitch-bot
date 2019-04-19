package main

import (
	"flag"
	"github.com/hueyjj/twitch-bot/client"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var addr = flag.String("addr", "irc-ws.chat.twitch.tv:443", "twitch irc-ws service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botUsername := os.Getenv("BOT_USERNAME")
	channelName := os.Getenv("CHANNEL_NAME")
	oAuthToken := os.Getenv("OAUTH_TOKEN")

	if botUsername == "" || channelName == "" || oAuthToken == "" {
		log.Fatalf("Environmental variables not set: BOT_USERNAME=%s CHANNEL_NAME=%s OAUTH_TOKEN=%s\n", botUsername, channelName, oAuthToken)
	} else {
		log.Printf("Environmental variables set: BOT_USERNAME=%s CHANNEL_NAME=%s OAUTH_TOKEN=%s\n", botUsername, channelName, oAuthToken)
	}

	client.Run(*addr, &client.Client{
		BotUsername: botUsername,
		ChannelName: channelName,
		OAuthToken:  oAuthToken,
		Conn:        nil,
	})
}
