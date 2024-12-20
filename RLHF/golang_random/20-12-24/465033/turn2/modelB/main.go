package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, workChan <-chan int) {
	defer wg.Done()
	for job := range workChan {
		fmt.Printf("Worker %d: Doing job %d\n", id, job)
		time.Sleep(time.Second * 1) // Simulate work time
	}
}
func main() {
	const numWorkers = 3
	const numJobs = 5
	var wg sync.WaitGroup
	workChan := make(chan int, numJobs)
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, workChan)
	}
	// Send jobs to the worker goroutines
	for i := 1; i <= numJobs; i++ {
		workChan <- i
	}
	close(workChan) // Close the channel to indicate no more jobs
	wg.Wait()       // Wait for all workers to complete their jobs
	fmt.Println("All jobs are complete.")
}
