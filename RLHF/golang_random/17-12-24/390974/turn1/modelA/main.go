package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers  = 3
	numJobs     = 10
	jobDuration = 2 * time.Second
)

// Job represents a unit of work
type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate work
		time.Sleep(jobDuration)
		result := fmt.Sprintf("Worker %d completed Job %d", id, job.ID)
		results <- result
	}
}

func fanOut(jobs chan Job, results chan string) {
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func fanIn(resultsChans []chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	output := func(c chan string) {
		for result := range c {
			out <- result
		}
		wg.Done()
	}

	for _, ch := range resultsChans {
		wg.Add(1)
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create job channel and result channels for each worker
	jobs := make(chan Job, numJobs)
	resultsChans := make([]chan string, numWorkers)

	for i := 0; i < numWorkers; i++ {
		resultsChans[i] = make(chan string)
	}

	// Start fan-out
	fanOut(jobs, resultsChans[0]) // Start the first worker as a reference for fan-out

	// Add jobs to the jobs channel
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- Job{ID: i}
		}
		close(jobs)
	}()

	// Use fan-in to aggregate results from all workers
	finalResults := fanIn(resultsChans)

	// Process final results
	for result := range finalResults {
		fmt.Println(result)
	}
}
