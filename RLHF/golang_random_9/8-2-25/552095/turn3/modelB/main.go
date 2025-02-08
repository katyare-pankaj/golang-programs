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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // You should customize this for security reasons
	},
}

type client struct {
	conn *websocket.Conn
	done chan struct{}
}

func (c *client) handleConnection() {
	defer c.conn.Close()
	defer close(c.done)

	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		log.Printf("Received message: %s", message)

		if err := c.conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Error writing message: %v", err)
			return
		}
	}
}

func newClient(conn *websocket.Conn) *client {
	return &client{
		conn: conn,
		done: make(chan struct{}),
	}
}

func main() {
	var wg sync.WaitGroup
	clientCh := make(chan *client)
	doneCh := make(chan struct{})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}
		client := newClient(conn)
		wg.Add(1)
		clientCh <- client
	})

	serverAddr := "localhost:8080"
	server := &http.Server{Addr: serverAddr}

	go func() {
		log.Printf("WebSocket server started on %s", serverAddr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	go func() {
		for {
			select {
			case client := <-clientCh:
				go client.handleConnection()
				wg.Done()
			case <-doneCh:
				// Server is shutting down, close all connections
				log.Println("Server is shutting down, closing connections")
				close(clientCh)
				return
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutdown signal received")
	close(doneCh)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}

	// Wait for all client connections to be handled
	wg.Wait()
	log.Println("All connections closed, exiting.")
}
