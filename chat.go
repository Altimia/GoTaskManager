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

func (c *Chat) SendMessage(message string) error {
	c.Messages = append(c.Messages, message)
	err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	fmt.Printf("%s sent a message: %s\n", c.From.Username, message)
	return nil
}

func (c *Chat) ReceiveMessage() error {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("error receiving message: %w", err)
		}
		c.Messages = append(c.Messages, string(message))
		fmt.Printf("%s received a message: %s\n", c.To.Username, string(message))
	}
}
