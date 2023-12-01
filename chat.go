package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// WebSocketConnection defines the interface for a websocket connection.
type WebSocketConnection interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
}

type Chat struct {
	ID       int
	Messages []string
	From     User
	To       User
	Conn     WebSocketConnection
}

func NewChat(id int, from User, to User, conn WebSocketConnection) *Chat {
	return &Chat{
		ID:       id,
		Messages: []string{},
		From:     from,
		To:       to,
		Conn:     conn,
	}
}

func (c *Chat) SendMessage(message string) error {
	if c.Conn == nil {
		return fmt.Errorf("no websocket connection")
	}
	c.Messages = append(c.Messages, message)
	if c.Conn != nil {
		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			zap.L().Error("Error sending message", zap.Error(err))
			return fmt.Errorf("error sending message: %w", err)
		}
	} else {
		zap.L().Error("Attempted to send message without a websocket connection")
		return fmt.Errorf("no websocket connection")
	}
	zap.L().Info("Message sent", zap.String("from", c.From.Username), zap.String("message", message))
	return nil
}

func (c *Chat) ReceiveMessage(stopChan <-chan struct{}) error {
	if c.Conn == nil {
		return fmt.Errorf("no websocket connection")
	}
	if c.Conn != nil {
		for {
			select {
			case <-stopChan:
				return nil
			default:
				_, message, err := c.Conn.ReadMessage()
				if err != nil {
					zap.L().Error("Error receiving message", zap.Error(err))
					return fmt.Errorf("error receiving message: %w", err)
				}
				c.Messages = append(c.Messages, string(message))
				zap.L().Info("Message received", zap.String("to", c.To.Username), zap.String("message", string(message)))
			}
		}
	} else {
		zap.L().Error("Attempted to receive message without a websocket connection")
		return fmt.Errorf("no websocket connection")
	}
}
