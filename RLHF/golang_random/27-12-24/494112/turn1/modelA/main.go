package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		// Simulate some work
		time.Sleep(time.Millisecond * 1)
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	runtime.GOMAXPROCS(4) // Set the maximum number of CPU cores
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go worker(i)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
