package main

import (
	"calculator/Connections"
	"log"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("http server started on :%v\n", port)
	err := http.ListenAndServe(":" +port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
