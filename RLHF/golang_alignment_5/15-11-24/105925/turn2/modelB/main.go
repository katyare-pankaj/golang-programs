package main

import (
	"fmt"
	"sync"
	"time"
)

var db map[string]string = make(map[string]string)
var wg sync.WaitGroup

func writeToDB(key, value string) {
	defer wg.Done()
	fmt.Println("Starting write:", key)
	time.Sleep(2 * time.Second) // Simulate write delay
	db[key] = value
	fmt.Println("Completed write:", key)
}

func readFromDB(key string) string {
	fmt.Println("Starting read:", key)
	time.Sleep(1 * time.Second) // Simulate read delay
	value, ok := db[key]
	if !ok {
		fmt.Println("Key not found:", key)
		return ""
	}
	fmt.Println("Completed read:", key, "->", value)
	return value
}

func main() {
	fmt.Println("Monolithic Database Service")
	wg.Add(3)

	go writeToDB("user1", "Alice")
	go writeToDB("user2", "Bob")
	go readFromDB("user1")

	wg.Wait()
	fmt.Println("Monolithic process completed.")
}
