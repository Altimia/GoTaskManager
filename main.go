package main

import (
	"context"
	"log"
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
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading GET request to websocket: %v", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := ws.Close(); err != nil {
			log.Printf("Error closing websocket: %v", err)
		}
	}() // Make sure we close the connection when the function returns

	// Register the connection with the user's ID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		log.Printf("Error converting user_id to int: %v", err)
		return
	}
	userConnections[uint(userID)] = ws
	log.Printf("User with id %d connected via websocket", userID)

	// Infinite loop to continuously read incoming messages
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(userConnections, uint(userID)) // Remove the connection if there's an error
			log.Printf("User with id %d disconnected", userID)
			return
		}
		// Log the message
		log.Printf("Message received from user with id %d: %s", userID, msg)
	}
}

func main() {
	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Create a channel to listen for interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	server := &http.Server{Addr: ":8000", Handler: nil}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	log.Println("Server started on :8000")

	// Block until a signal is received
	<-signalChan
	log.Println("Shutting down server...")

	// Call CloseAPI to handle graceful shutdown of the API server
	CloseAPI()
	CloseAPI()

	// Optionally, you can add a timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Println("Server exited properly")
}
