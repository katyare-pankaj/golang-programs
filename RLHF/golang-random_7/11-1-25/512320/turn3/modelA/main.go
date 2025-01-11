package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := []string{"job1", "job2", "job3", "job4", "job5"}
	processJobsConcurrently(jobs)
}

func processJobsConcurrently(jobs []string) {
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	// Ensure each goroutine gets its own copy of the job name
	for _, job := range jobs {
		go func(job string) {
			defer wg.Done()
			processJob(job)
		}(job)
	}

	wg.Wait()
	fmt.Println("All jobs processed.")
}

func processJob(job string) {
	// Simulate some work for the job
	fmt.Println("Processing job:", job)

	// Simulate processing duration (e.g., 1 second)
	time.Sleep(time.Second)
}
