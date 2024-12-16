package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function that sends data to the channel
func producer(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		value := rand.Intn(100)
		fmt.Printf("Producer sending: %d\n", value)
		ch <- value
		time.Sleep(time.Millisecond * 50) // simulate work
	}
}

// Consumer function that receives data from the channel
func consumer(ch <-chan int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		value, ok := <-ch
		if !ok {
			fmt.Println("Consumer received end signal.")
			return
		}
		fmt.Printf("Consumer received: %d\n", value)
		time.Sleep(time.Millisecond * 200) // simulate work
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const workCount = 100
	var wg sync.WaitGroup

	fmt.Println("Unbuffered Channel Example:")
	// Unbuffered channel
	unbufferedCh := make(chan int)

	wg.Add(1)
	go producer(unbufferedCh, workCount, &wg)

	wg.Add(1)
	go consumer(unbufferedCh, workCount, &wg)

	wg.Wait()
	close(unbufferedCh)

	fmt.Println("\nBuffered Channel Example (buffer size = 10):")
	// Buffered channel with a size of 10
	bufferedCh := make(chan int, 10)

	wg.Add(1)
	go producer(bufferedCh, workCount, &wg)

	wg.Add(1)
	go consumer(bufferedCh, workCount, &wg)

	wg.Wait()
	close(bufferedCh)
}
