package client

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

var (
	// PongMessage about every five minutes, twitch sends a "PING :tmi.twitch.tv"
	// we return a pong message to keep our connection alive.
	PongMessage = "PONG :tmi.twitch.tv"

	// PrivMsgRe matches messages sent by users in a channel
	// NOTE: Maybe need to validate message by channel instead `.+`.
	PrivMsgRe = regexp.MustCompile(`PRIVMSG #.+ :`)

	// UsernameTmiRe matches strings/messages for a UsernameRe
	// e.g. ...ot@muddycoacoabot.tmi.twitch.tv PRIV... matches @muddycoacoabot.tmi.twitch.tv
	UsernameTmiRe = regexp.MustCompile(`@.+[.]tmi[.]twitch[.]tv`)

	// UsernameRe matches any strings , used to find the first valid string in
	// @muddycoacoabot.tmi.twitch.tv, for example.
	UsernameRe = regexp.MustCompile(`[^@][^.]+`)
)

// AuthMessage returns a formated authentication string
// e.g. PASS oauth:0123456789abcdefghijABCDEFGHIJ
func AuthMessage(oAuthToken string) string {
	return fmt.Sprintf("PASS %s", oAuthToken)
}

// NickMessage returns a formatted nickname string to be authenticated as
// e.g. NICK jojojostar
func NickMessage(nick string) string {
	return fmt.Sprintf("NICK %s", nick)
}

// JoinMessage returns a formated channel string to join
// e.g. JOIN #shroud
func JoinMessage(channel string) string {
	return fmt.Sprintf("JOIN #%s", channel)
}

// ChannelMessage returns a formatted string to send to the channel
// e.g PRIVMSG #shroud :I love coconuts!
func ChannelMessage(channel, msg string) string {
	return fmt.Sprintf("PRIVMSG #%s :%s", channel, msg)
}

// ParseUsername looks for the first valid string in @myusername.tmi.twitch.tv
// e.g. myusername
func ParseUsername(msg string) (string, error) {
	msgBytes := []byte(msg)
	if match := UsernameTmiRe.Find(msgBytes); match != nil {
		log.Println("match=", string(match))
		if username := UsernameRe.Find(match); username != nil {
			return string(username), nil
		}
	}
	return "", errors.New("Unable to parse username")
}
