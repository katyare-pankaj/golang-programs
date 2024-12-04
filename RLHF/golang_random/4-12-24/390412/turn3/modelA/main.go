package main

import (
	"fmt"
	"sync"
	"time"
)

func work(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		time.Sleep(time.Millisecond * time.Duration(i*10))
		fmt.Println("Processed:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	workChannel := make(chan int, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go work(workChannel, &wg)
	}

	for i := 0; i < 100; i++ {
		workChannel <- i
	}

	close(workChannel)
	wg.Wait()
}
