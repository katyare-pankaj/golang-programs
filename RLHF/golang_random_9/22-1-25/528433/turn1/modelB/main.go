package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	workCount := 200

	start := time.Now()
	for i := 0; i < workCount; i++ {
		wg.Add(1)
		go doWork(i, &wg)
	}
	wg.Wait()
	end := time.Now()

	fmt.Printf("Total time taken: %s\n", end.Sub(start))
	fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
}

func doWork(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Millisecond) // Simulate work
	fmt.Printf("Worker %d completed\n", id)
}
