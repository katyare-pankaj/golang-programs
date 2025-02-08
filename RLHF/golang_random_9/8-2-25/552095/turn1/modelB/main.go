package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Define the WaitGroup and initialize it with the desired number of goroutines (in this case, 2).
var wg sync.WaitGroup

const numGoroutines = 2

func init() {
	wg.Add(numGoroutines)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a websocket.
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer c.Close()

	// Start a new goroutine to handle the connection, ensuring that the connection is closed after the goroutine finishes.
	go func() {
		defer wg.Done()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			log.Printf("Received message: %s", message)
			// Process the received message here...

			// Send back a pong message to the client
			err = c.WriteMessage(mt, []byte("pong"))
			if err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	}()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", handleConnection)
	fmt.Println("HTTP server listening on port 8080")
	// Start the http server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	// Wait for all goroutines to complete before exiting the program.
	wg.Wait()
	fmt.Println("All connections closed. Exiting.")
}
