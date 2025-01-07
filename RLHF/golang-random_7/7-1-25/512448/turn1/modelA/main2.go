package main

import (
	"fmt"
	"sync"
	"time"
)

func simulateDataStream(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(time.Second) // Simulate data arrival delay
		fmt.Println("Stream", id, ": New data arrived")
	}
}

func main() {
	var wg sync.WaitGroup
	numStreams := 3

	for i := 0; i < numStreams; i++ {
		wg.Add(1)
		go simulateDataStream(i, &wg)
	}

	select {
	case <-time.After(time.Duration(1) * time.Second):
		fmt.Println("Stopping data streams after 5 seconds.")
		wg.Wait()
		fmt.Println("All streams stopped.")
	}
}
