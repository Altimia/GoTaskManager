package main

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	userConnections = make(map[uint]*websocket.Conn)
	userConnectionsMutex sync.Mutex
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		zap.L().Error("Error upgrading GET request to websocket", zap.Error(err))
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := ws.Close(); err != nil {
			zap.L().Error("Error closing websocket", zap.Error(err))
		}
	}() // Make sure we close the connection when the function returns

	// Register the connection with the user's ID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		zap.L().Error("Error converting user_id to int", zap.Error(err))
		return
	}
	userConnectionsMutex.Lock()
	userConnections[uint(userID)] = ws
	userConnectionsMutex.Unlock()
	zap.L().Info("User connected via websocket", zap.Int("userID", userID))

	// Infinite loop to continuously read incoming messages
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			zap.L().Error("Error reading message", zap.Error(err))
			userConnectionsMutex.Lock()
			delete(userConnections, uint(userID)) // Remove the connection if there's an error
			userConnectionsMutex.Unlock()
			zap.L().Info("User disconnected", zap.Int("userID", userID))
			return
		}
		// Log the message
		zap.L().Info("Message received from user", zap.Int("userID", userID), zap.ByteString("message", msg))
	}
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	zap.ReplaceGlobals(logger)
	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Create a channel to listen for interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	server := &http.Server{Addr: ":8000", Handler: nil}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("Error starting server", zap.Error(err))
		}
	}()
	zap.L().Info("Server started", zap.String("address", ":8000"))

	// Block until a signal is received
	<-signalChan
	zap.L().Info("Shutting down server...")

	// Call CloseAPI to handle graceful shutdown of the API server
	CloseAPI()
	CloseAPI()

	// Optionally, you can add a timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown failed", zap.Error(err))
	}
	zap.L().Info("Server exited properly")
}
