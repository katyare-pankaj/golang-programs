package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// upgrader is used to upgrade an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true

	},
}

func handleConnection(conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	defer conn.Close()

	for {
		// ReadMessage is a helper method for getting a complete message as a slice of bytes
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message: %s", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	wg.Add(1) // Increment the WaitGroup counter before starting a new goroutine
	go handleConnection(conn, wg)
}

func main() {
	var wg sync.WaitGroup

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocketHandler(w, r, &wg)
	})

	serverAddr := "localhost:8080"
	fmt.Printf("WebSocket server started at %s\n", serverAddr)

	go func() {
		log.Fatal(http.ListenAndServe(serverAddr, nil))
	}()

	// Wait for all the goroutines to finish before exiting
	wg.Wait()
	fmt.Println("All connections have been handled.")
}
