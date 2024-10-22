package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hashicorp/consul/api"
)

const (
	lockKey           = "my-distributed-lock"
	lockHoldTimeout   = 5 * time.Second
	lockRetainTimeout = 10 * time.Second
)

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Consul client configuration
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	session := client.Session()
	sessionID, _, err := session.Create(&api.SessionEntry{
		TTL:       lockRetainTimeout.String(),
		LockDelay: lockHoldTimeout,
	}, nil)

	defer session.Destroy(sessionID, nil)

	for {
		if acquireLock(client, sessionID) {
			fmt.Println("Lock acquired!")
			// Simulate work being done with the lock
			doWork()
			releaseLock(client, sessionID)
			fmt.Println("Lock released.")
		} else {
			fmt.Println("Lock not acquired, retrying...")
		}
		// Add random delay before retrying
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func acquireLock(client *api.Client, sessionID string) bool {
	opts := &api.WriteOptions{}

	for {
		success, _, err := client.KV().Acquire(&api.KVPair{Key: lockKey}, opts)
		if err != nil {
			fmt.Println("Error acquiring lock:", err)
			return false
		}

		if success {
			return true
		}

		// Add random delay before retrying
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func releaseLock(client *api.Client, sessionID string) {
	opts := &api.WriteOptions{}
	fmt.Println("Writer :", opts)
}

func doWork() {
	// Simulate some work being done with the lock held
	time.Sleep(2 * time.Second)
}
