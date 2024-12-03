package main

import (
	"fmt"
)

func main() {
	count := 10
	ch := make(chan int)

	// Goroutine to send values to channel
	go func() {
		for i := 0; i < count; i++ {
			ch <- i // Send data to channel
		}
		close(ch) // Close channel when done
	}()

	// Reading from the channel
	for num := range ch {
		fmt.Println(num)
	}
}
