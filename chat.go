package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Chat struct {
	ID       int
	Messages []string
	From     User
	To       User
	Conn     *websocket.Conn
}

func NewChat(id int, from User, to User, conn *websocket.Conn) *Chat {
	return &Chat{
		ID:       id,
		Messages: []string{},
		From:     from,
		To:       to,
		Conn:     conn,
	}
}

func (c *Chat) SendMessage(message string) error {
	c.Messages = append(c.Messages, message)
	if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		zap.L().Error("Error sending message", zap.Error(err))
		return fmt.Errorf("Error sending message: %w", err)
	}
	zap.L().Info("Message sent", zap.String("from", c.From.Username), zap.String("message", message))
	return nil
}

func (c *Chat) ReceiveMessage() error {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			zap.L().Error("Error receiving message", zap.Error(err))
			return fmt.Errorf("Error receiving message: %w", err)
		}
		c.Messages = append(c.Messages, string(message))
		zap.L().Info("Message received", zap.String("to", c.To.Username), zap.String("message", string(message)))
	}
}
