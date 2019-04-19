package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"time"
)

var addr = flag.String("addr", "irc-ws.chat.twitch.tv:443", "twitch irc-ws service address")

func main() {
	flag.Parse()

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

	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	re := regexp.MustCompile(`[#].*:`)
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("> %s", message)

			if match := re.Find([]byte(message)); match != nil {
				username := string(match)
				username = username[1 : len(username)-1]
				log.Printf("username=%s\n", username)
				respMsg := fmt.Sprintf("PRIVMSG #%s :Your waifu is trash. %s", channelName, username)
				log.Printf("< %s\n", respMsg)
				err = c.WriteMessage(websocket.TextMessage, []byte(respMsg))
				if err != nil {
					log.Println("Auth msg:", err)
				}
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	authMsg := fmt.Sprintf("PASS %s", oAuthToken)
	log.Printf("< %s\n", authMsg)
	err = c.WriteMessage(websocket.TextMessage, []byte(authMsg))
	if err != nil {
		log.Println("Auth msg:", err)
	}

	nickMsg := fmt.Sprintf("NICK %s", botUsername)
	log.Printf("< %s\n", nickMsg)
	err = c.WriteMessage(websocket.TextMessage, []byte(nickMsg))
	if err != nil {
		log.Println("Nick message:", err)
	}

	joinChannelMsg := fmt.Sprintf("JOIN #%s", channelName)
	log.Printf("< %s\n", joinChannelMsg)
	err = c.WriteMessage(websocket.TextMessage, []byte(joinChannelMsg))
	if err != nil {
		log.Println("Join channel message:", err)
	}

	for {
		select {
		case <-done:
			return
		//case t := <-ticker.C:

		//err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		//if err != nil {
		//	log.Println("write:", err)
		//	return
		//}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
