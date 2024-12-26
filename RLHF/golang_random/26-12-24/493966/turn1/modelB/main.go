package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d: %d\n", id, i)
		// Simulate work by introducing a short delay
		//time.Sleep(time.Duration(100 * rand.Intn(5)) * time.Millisecond)
	}
}
func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers completed.")
}
