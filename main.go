package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		log.Printf("error upgrading GET request to websocket: %v", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer ws.Close() // Make sure we close the connection when the function returns

	// Register the connection with the user's ID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		log.Printf("error converting user_id to int: %v", err)
		return
	}
	userConnections[uint(userID)] = ws

	// Infinite loop to continuously read incoming messages
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			delete(userConnections, uint(userID)) // Remove the connection if there's an error
			return
		}
		// Print the message to the console
		fmt.Printf("%s\n", msg)
	}
}

import (
	"os"
	"os/signal"
	"syscall"
)

// ...

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
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Block until a signal is received
	<-signalChan
	fmt.Println("Shutting down server...")

	// Call CloseAPI to handle graceful shutdown
	CloseAPI()

	// Optionally, you can add a timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	fmt.Println("Server exited properly")
}
