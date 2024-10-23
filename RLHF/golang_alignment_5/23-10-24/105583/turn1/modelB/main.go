package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
)

const (
	natsURL        = "nats://localhost:4222"
	redisURL       = "redis://localhost:6379"
	socialMediaURL = "https://api.example.com"
	topic          = "social.events"
)

var (
	nc       *nats.Conn
	rdb      *redis.Client
	httpAddr = ":8080"
)

func main() {
	// Set up Redis and NATS connections
	rdb = redis.NewClient(&redis.Options{Addr: redisURL})
	defer rdb.Close()

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	// Subscribe to the NATS topic for social media events
	_, err = nc.Subscribe(topic, func(m *nats.Msg) {
		handleEvent(m)
	})
	if err != nil {
		log.Fatalf("Error subscribing to NATS topic: %v", err)
	}

	// Start the HTTP server
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	go func() {
		log.Fatal(http.ListenAndServe(httpAddr, r))
	}()

	// Graceful shutdown handling
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	fmt.Println("Shutting down...")
}

func handleEvent(m *nats.Msg) {
	// Parse the event message and perform actions based on the event type
	// For simplicity, we'll just log the event message to Redis.
	err := rdb.Set(context.Background(), "social_event", string(m.Data), 0).Err()
	if err != nil {
		log.Printf("Error handling event: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
