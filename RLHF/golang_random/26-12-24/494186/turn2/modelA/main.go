package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done() is called even if the function panics
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate work
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3) // Start with 3 workers

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	fmt.Println("Main starting other tasks...")
	time.Sleep(time.Second)
	fmt.Println("Main waiting for workers...")
	wg.Wait()
	fmt.Println("All workers done, main continuing...")
}
