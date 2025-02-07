package main

import (
	"fmt"
	"sync"
)

const (
	poolSize   = 5  // Number of workers
	numWorkers = 10 // Total number of tasks
)

// SharedSlice encapsulates a slice with a mutex for safe concurrent access.
type SharedSlice struct {
	sync.Mutex
	data []int
}

// Add adds a value to the slice safely.
func (s *SharedSlice) Add(value int) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, value)
}

// GetData safely returns a copy of the displayed slice's data.
func (s *SharedSlice) GetData() []int {
	s.Lock()
	defer s.Unlock()
	copyData := make([]int, len(s.data))
	copy(copyData, s.data)
	return copyData
}

func worker(id int, tasks <-chan int, sharedSlice *SharedSlice, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		sharedSlice.Add(task)
	}
}

func main() {
	tasks := make(chan int, numWorkers) // Task queue
	var wg sync.WaitGroup

	sharedSlice := &SharedSlice{data: []int{}}

	// Start workers
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(i, tasks, sharedSlice, &wg)
	}

	// Dispatch tasks
	for i := 0; i < numWorkers; i++ {
		tasks <- i
	}
	close(tasks) // Signals all tasks have been sent

	// Wait for all workers to finish
	wg.Wait()

	// Output the results
	fmt.Println("Final slice contents:", sharedSlice.GetData())
}
