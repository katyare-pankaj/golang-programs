package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function that sends data to the channel
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		value := rand.Intn(100)
		fmt.Printf("Producer %d produced: %d\n", id, value)
		ch <- value
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	}
}

// Consumer function that receives data from the channel
func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("Consumer %d consumed: %d\n", id, value)
		time.Sleep(time.Millisecond * 150) // simulate processing
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Small buffer size
	smallBuffer := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(1, smallBuffer, &wg)

	wg.Add(1)
	go consumer(1, smallBuffer, &wg)

	wg.Wait()
	close(smallBuffer)

	fmt.Println("Small buffer processing done.")

	// Large buffer size
	largeBuffer := make(chan int, 10)
	wg.Add(1)
	go producer(2, largeBuffer, &wg)

	wg.Add(1)
	go consumer(2, largeBuffer, &wg)

	wg.Wait()
	close(largeBuffer)

	fmt.Println("Large buffer processing done.")
}
