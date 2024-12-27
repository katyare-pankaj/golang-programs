package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg, 1)
	go worker(&wg, 2)

	fmt.Println("Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("All workers finished.")
}
