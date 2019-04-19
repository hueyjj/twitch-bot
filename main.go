package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
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
}
