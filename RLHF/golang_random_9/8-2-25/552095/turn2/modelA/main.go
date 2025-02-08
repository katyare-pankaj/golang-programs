package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader with an origin check
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handleConnection deals with WebSocket messages
func handleConnection(conn *websocket.Conn, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()
	defer conn.Close()

	for {
		select {
		case <-done:
			// Server is shutting down
			log.Println("Server is shutting down, closing connection")
			return
		default:
			// Read message from client
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				return
			}
			log.Printf("Received message: %s", message)

			// Echo the message back to the client
			if err := conn.WriteMessage(messageType, message); err != nil {
				log.Printf("Error writing message: %v", err)
				return
			}
		}
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup, done chan struct{}) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	wg.Add(1)
	go handleConnection(conn, wg, done)
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocketHandler(w, r, &wg, done)
	})

	// Setup server
	serverAddr := "localhost:8080"
	server := &http.Server{Addr: serverAddr}

	// Shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("WebSocket server started on %s", serverAddr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-stop
	log.Println("Shutdown signal received")

	// Signal all goroutines to finish work
	close(done)

	// Create a context with a timeout to give outstanding connections time to finish
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Initiate graceful shutdown of the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	// Wait for all websocket connections to be handled
	wg.Wait()
	log.Println("All connections closed, exiting.")
}
