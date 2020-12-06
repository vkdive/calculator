package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan []byte)           // broadcast channel
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../frontEnd"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/update", handleConnections)

	// Start listening for incoming messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	// Register our new client
	clients[ws] = true
	for {
		var p []byte
		_, p, err = ws.ReadMessage()
		if err != nil {
			log.Printf("error while reading from client: %v", err)
			delete(clients, ws)
			break
		}
		msg:= bytes.TrimSpace(bytes.Replace(p, newline, space, -1))

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("error whule writing to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
