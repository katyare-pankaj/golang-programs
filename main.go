package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Publish a message
	message := "Hello, NATS!"
	err = nc.Publish("test.subject", []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Published message:", message)
}
