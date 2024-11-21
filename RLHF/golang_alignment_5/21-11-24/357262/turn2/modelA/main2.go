package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered channel with size 3
	ch := make(chan int, 3)

	go func() {
		fmt.Println("Goroutine 1 starting...")
		for i := 1; i <= 5; i++ {
			ch <- i // Sends 5 messages, buffered if receiver is not ready
			fmt.Println("Goroutine 1 sent:", i)
			time.Sleep(time.Millisecond * 50)
		}
		fmt.Println("Goroutine 1 finished sending.")
	}()

	go func() {
		fmt.Println("Goroutine 2 starting...")
		for msg := range ch {
			fmt.Println("Goroutine 2 received:", msg)
			time.Sleep(time.Millisecond * 100)
		}
		fmt.Println("Goroutine 2 finished.")
	}()

	fmt.Println("Main goroutine starting...")
	time.Sleep(time.Second)
	fmt.Println("Main goroutine finished.")
}
