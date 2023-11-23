package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Infinite loop to continuously read incoming messages
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		// Print the message to the console
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start the server on localhost port 8000 and log any errors
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
