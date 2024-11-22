package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		data := rand.Intn(100)
		ch <- data // Send data to the channel
		fmt.Printf("Producer %d sent: %d\n", id, data)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Consumer processed: %d\n", data)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	// Create a buffered channel with capacity 10
	ch := make(chan int, 10)

	// Start 3 producers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go producer(i, ch, &wg)
	}

	// Start 1 consumer
	wg.Add(1)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("All done!")
}
