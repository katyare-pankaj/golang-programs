package main

import (
	"fmt"
	"sync"
	"time"
)

func slowReader(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range c {
		// Simulate slow reading operation
		time.Sleep(time.Second * 2)
		fmt.Println("Read value:", val)
	}
}

func writer(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	// Launch a single slow reader
	wg.Add(1)
	go slowReader(c, &wg)

	// Start writing to the channel
	go writer(c)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
