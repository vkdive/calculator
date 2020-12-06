package main

import (
	"calculator/Connections"
	"log"
	"net/http"
)


func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("./frontEnd"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/update", Connections.HandleConnections)

	// Start listening for incoming messages
	go Connections.HandleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
