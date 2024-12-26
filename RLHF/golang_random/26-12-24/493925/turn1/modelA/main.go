package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Worker", id, "starting...")
	// Simulate work
	for i := 0; i < 10; i++ {
		fmt.Println("Worker", id, "doing work...")
	}
	fmt.Println("Worker", id, "finished.")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	fmt.Println("Main starting to wait...")
	wg.Wait()
	fmt.Println("Main finished waiting.")
}
