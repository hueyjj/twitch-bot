package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var ()

// Client holds authentication, bot username, and which channel to send message to
type Client struct {
	BotUsername string
	ChannelName string
	OAuthToken  string
	Conn        *websocket.Conn
}

// Run make a websocket client to read from the twitch server and write back
func Run(serverAddr string, client *Client) {
	u := url.URL{Scheme: "wss", Host: serverAddr}
	log.Printf("Connecting to %s", u.String())

	// Dial the server and establish a connection
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	client.Conn = conn

	// Receive messages from the twitch server
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read message:", err)
				return
			}
			msg := string(message)
			log.Println(">", msg)

			username, err := ParseUsername(msg)
			if err != nil {
				log.Println("ParseUsername:", err)
			}
			log.Println("username=", username)

			//if match := re.Find([]byte(message)); match != nil {
			//	username := string(match)
			//	username = username[1 : len(username)-1]
			//	log.Printf("username=%s\n", username)
			//	//respMsg := fmt.Sprintf("PRIVMSG #%s :Your waifu is trash. %s", channelName, username)
			//	log.Printf("< %s\n", respMsg)
			//	err = c.WriteMessage(websocket.TextMessage, []byte(ChannelMessage(channelName, username)))
			//	if err != nil {
			//		log.Println("Auth msg:", err)
			//	}
			//}
		}
	}()

	SendMessage(AuthMessage(client.OAuthToken), client)
	SendMessage(NickMessage(client.BotUsername), client)
	SendMessage(JoinMessage(client.ChannelName), client)

	//authMsg := fmt.Sprintf("PASS %s", oAuthToken)
	//log.Printf("< %s\n", authMsg)
	//err = c.WriteMessage(websocket.TextMessage, []byte(authMsg))
	//if err != nil {
	//	log.Println("Auth msg:", err)
	//}

	//nickMsg := fmt.Sprintf("NICK %s", botUsername)
	//log.Printf("< %s\n", nickMsg)
	//err = c.WriteMessage(websocket.TextMessage, []byte(nickMsg))
	//if err != nil {
	//	log.Println("Nick message:", err)
	//}

	//joinChannelMsg := fmt.Sprintf("JOIN #%s", channelName)
	//log.Printf("< %s\n", joinChannelMsg)
	//err = c.WriteMessage(websocket.TextMessage, []byte(joinChannelMsg))
	//if err != nil {
	//	log.Println("Join channel message:", err)
	//}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			// Every minute send a PONG message to keep connection alive
			SendMessage(PongMessage, client)
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
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

// SendMessage sends a message to the connection in the client
func SendMessage(msg string, client *Client) {
	log.Printf("< %s\n", msg)
	err := client.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("send message:", err)
	}
}
