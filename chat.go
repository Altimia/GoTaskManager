package main

import (
	"fmt"
)

type Chat struct {
	ID       int
	Messages []string
	From     User
	To       User
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
	fmt.Printf("%s sent a message: %s\n", c.From.Username, message)
}

func (c *Chat) ReceiveMessage(message string) {
	c.Messages = append(c.Messages, message)
	fmt.Printf("%s received a message: %s\n", c.To.Username, message)
}
