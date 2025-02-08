package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Use your own origin validation logic
		},
	}
	// Store client connections in a map with client IDs
	clients     = make(map[string]*websocket.Conn)
	clientMutex sync.Mutex
	// WaitGroup to track active goroutines
	wg sync.WaitGroup
)

func handleConnection(clientID string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client %s: Error reading message: %v", clientID, err)
			break
		}
		log.Printf("Client %s: Received message: %s", clientID, message)
		// Echo message back to all clients
		broadcastMessage(message)
	}
	// Remove client from map when connection is closed
	clientMutex.Lock()
	delete(clients, clientID)
	clientMutex.Unlock()
	log.Printf("Client %s disconnected.", clientID)
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}

	// Generate a unique client ID
	clientID := generateClientID()
	// Add the new client to the map
	clientMutex.Lock()
	clients[clientID] = conn
	clientMutex.Unlock()

	wg.Add(1)
	go handleConnection(clientID, conn, &wg)
	log.Printf("Client %s connected.", clientID)
}

// Function to generate a unique client ID (you can use better methods like UUIDs)
func generateClientID() string {
	return fmt.Sprintf("client-%d", len(clients))
}

// Function to broadcast message to all connected clients
func broadcastMessage(message []byte) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	for clientID, conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Client %s: Error writing message: %v", clientID, err)
			conn.Close()
			delete(clients, clientID)
		}
	}
}

func main() {
	http.HandleFunc("/ws", websocketHandler)

	serverAddr := ":8080"
	fmt.Printf("WebSocket server started at %s\n", serverAddr)

	go func() {
		err := http.ListenAndServe(serverAddr, nil)
		if err != nil {
			log.Fatal("Error starting server:", err)
		}
	}()

	// Wait for all goroutines to finish before exiting
	wg.Wait()
	fmt.Println("All clients disconnected.")
}
