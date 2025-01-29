package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWithDeferInLoop(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		// Some actual work is done here...
	}
}

func workerWithDeferOutsideLoop(id int, wg *sync.WaitGroup) {
	cleanup := func() {
		wg.Done()
	}
	defer cleanup()
	for i := 0; i < 100000; i++ {
		// Some actual work is done here...
	}
}

func main() {
	const numWorkers = 10
	var wg sync.WaitGroup

	start := time.Now()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go workerWithDeferInLoop(i, &wg)
	}
	wg.Wait()
	elapsed1 := time.Since(start)
	fmt.Println("Elapsed time with defer in loop:", elapsed1)

	start = time.Now()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go workerWithDeferOutsideLoop(i, &wg)
	}
	wg.Wait()
	elapsed2 := time.Since(start)
	fmt.Println("Elapsed time with defer outside loop:", elapsed2)
}
