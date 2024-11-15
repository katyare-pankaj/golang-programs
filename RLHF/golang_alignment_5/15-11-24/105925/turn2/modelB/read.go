package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeToDB(key, value string) {
	defer wg.Done()
	fmt.Println("Starting write:", key)
	time.Sleep(2 * time.Second) // Simulate write delay
	db[key] = value
	fmt.Println("Completed write:", key)
}

func main() {
	fmt.Println("Write Service")
	wg.Add(2)

	go writeToDB("user1", "Alice")
	go writeToDB("user2", "Bob")

	wg.Wait()
	fmt.Println("Write process completed.")
}
