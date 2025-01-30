package main

import (
	"time"
)

func main() {
	// Mock data source; in a real system, this might be a socket or a queue
	dataStream := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			dataStream <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(dataStream)
	}()
}
