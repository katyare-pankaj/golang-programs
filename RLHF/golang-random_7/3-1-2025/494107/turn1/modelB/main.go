package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// Simulate work by sleeping for a random duration
	// sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	// time.Sleep(sleepDuration)
	fmt.Printf("Worker %d finished\n", id)
}
func main() {
	var wg sync.WaitGroup
	numWorkers := 5
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers completed.")
}
