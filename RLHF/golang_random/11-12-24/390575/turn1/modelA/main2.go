package main

import (
	"fmt"
	"time"
)

func increment(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- 1
	}
}

func main() {
	startTime := time.Now()

	ch := make(chan int, 10000) // buffered channel to avoid blocking
	var counter int

	for i := 0; i < 10; i++ {
		go increment(ch)
	}

	go func() {
		for i := 0; i < 10*1000; i++ {
			counter += <-ch
		}
	}()

	time.Sleep(1 * time.Second) // wait for goroutines to finish
	fmt.Printf("Channel Counter: %d, Time: %v\n", counter, time.Since(startTime))
}
