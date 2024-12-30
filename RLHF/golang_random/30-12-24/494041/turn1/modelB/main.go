package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Every goroutine must call Done to tell the WaitGroup it's finished
	for i := 0; i < 1000; i++ {
		fmt.Printf("Worker %d: %d\n", id, i)
	}
}
func main() {
	var wg sync.WaitGroup
	numGoroutines := 5
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(i, &wg)
	}
	wg.Wait() // Block the main goroutine until all worker goroutines have finished
	fmt.Println("All workers have completed their tasks.")
}
