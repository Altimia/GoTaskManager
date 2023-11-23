package main

import (
	"fmt"
)

type Chat struct {
	ID       int
	Messages []string
	From     User
	To       User
	Conn     *websocket.Conn
}

func NewChat(id int, message string, from User, to User) *Chat {
	return &Chat{
		ID:      id,
		Message: message,
		From:    from,
		To:      to,
	}
}

func (c *Chat) SendMessage(message string) {
	c.Messages = append(c.Messages, message)
	c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	fmt.Printf("%s sent a message: %s\n", c.From.Username, message)
}

func (c *Chat) ReceiveMessage() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		c.Messages = append(c.Messages, string(message))
		fmt.Printf("%s received a message: %s\n", c.To.Username, string(message))
	}
}
