package main

import (
	"errors"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWebSocketConn is a mock for the websocket.Conn
type MockWebSocketConn struct {
	mock.Mock
}

func (m *MockWebSocketConn) WriteMessage(messageType int, data []byte) error {
	args := m.Called(messageType, data)
	return args.Error(0)
}

func (m *MockWebSocketConn) ReadMessage() (messageType int, p []byte, err error) {
	args := m.Called()
	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func TestSendMessage(t *testing.T) {
	mockConn := new(MockWebSocketConn)
	mockConn.On("WriteMessage", websocket.TextMessage, []byte("test message")).Return(nil)

	chat := NewChat(1, User{Username: "fromUser"}, User{Username: "toUser"}, mockConn)
	err := chat.SendMessage("test message")

	assert.NoError(t, err)
	assert.Equal(t, []string{"test message"}, chat.Messages)
	mockConn.AssertExpectations(t)
}

func TestSendMessageNoConnection(t *testing.T) {
	chat := NewChat(1, User{Username: "fromUser"}, User{Username: "toUser"}, nil)
	err := chat.SendMessage("test message")

	assert.Error(t, err)
	assert.Equal(t, errors.New("no websocket connection"), err)
}

func TestReceiveMessage(t *testing.T) {
	mockConn := new(MockWebSocketConn)
	mockConn.On("ReadMessage").Return(websocket.TextMessage, []byte("received message"), nil).Once()
	mockConn.On("ReadMessage").Return(0, nil, errors.New("connection closed")).Once()

	chat := NewChat(1, User{Username: "fromUser"}, User{Username: "toUser"}, mockConn)
	stopChan := make(chan struct{})
	go func() {
		defer close(stopChan)
		chat.ReceiveMessage(stopChan)
	}()

	// Simulate receiving a message by sending a stop signal after a short delay
	time.AfterFunc(10*time.Millisecond, func() {
		stopChan <- struct{}{}
	})
	<-stopChan // Wait for the goroutine to finish

	assert.Equal(t, []string{"received message"}, chat.Messages)
	mockConn.AssertExpectations(t)
}

func TestReceiveMessageNoConnection(t *testing.T) {
	chat := NewChat(1, User{Username: "fromUser"}, User{Username: "toUser"}, nil)
	stopChan := make(chan struct{})
	err := chat.ReceiveMessage(stopChan)

	assert.Error(t, err)
	assert.Equal(t, errors.New("no websocket connection"), err)
}
