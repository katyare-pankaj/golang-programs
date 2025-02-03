package main

import (
	"fmt"
	"sync"
)

var counter int

func incrementCounter(wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		counter++
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Two goroutines will be running

	go incrementCounter(&wg)
	go incrementCounter(&wg)

	wg.Wait() // Wait for both goroutines to finish

	fmt.Println("Final Counter Value:", counter)
}
