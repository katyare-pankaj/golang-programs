package main

import (
	"fmt"
	"sync"
)

func producer(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func consumer(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for data := range c {
		fmt.Println("Received:", data)
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(1)
	go producer(&wg, c)

	wg.Add(1)
	go consumer(&wg, c)

	wg.Wait()
	// The program will hang here because the channel is not closed
}
