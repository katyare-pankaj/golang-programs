package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	jobs := []int{1, 2, 3, 4, 5}

	// Start 3 worker goroutines
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go worker(i+1, &jobs, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers have finished.")
	// All workers have processed the jobs. Now we can continue with the next task.
}

func worker(id int, jobs *[]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for len(*jobs) > 0 {
		job := popJob(jobs)
		fmt.Printf("Worker %d is processing job %d...\n", id, job)
		time.Sleep(time.Duration(job) * time.Second) // Simulate work
	}

	fmt.Printf("Worker %d has finished all jobs.\n", id)
}

// Helper function to pop a job from the job slice
func popJob(jobs *[]int) int {
	if len(*jobs) == 0 {
		return 0
	}

	job := (*jobs)[0]
	*jobs = (*jobs)[1:]
	return job
}
