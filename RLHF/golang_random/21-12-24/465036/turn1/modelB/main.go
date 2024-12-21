package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("-------------------")
	fmt.Println("Logistics with Goroutines and WaitGroup")
	fmt.Println("-------------------")
	// define workers and tasks
	workers := 3
	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5"}

	// traditional threading method
	start := time.Now()
	doTraditinalLogistics(workers, tasks)
	fmt.Println("Traditional Threading execution time:", time.Since(start))

	fmt.Println("\n\n-------------------")
	fmt.Println("Logistics with Goroutines and WaitGroup")
	fmt.Println("-------------------")

	// goroutine and waitgroup optimization
	start = time.Now()
	doLogisticsWithGoroutinesAndWaitGroup(workers, tasks)
	fmt.Println("Goroutine and WaitGroup execution time:", time.Since(start))

}

// traditional threading method
func doTraditinalLogistics(workers int, tasks []string) {
	// worker function
	worker := func(task string) {
		fmt.Println("Worker processing:", task)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker completed:", task)
	}

	// Calculate the task workload per worker
	tasksPerWorker := len(tasks) / workers
	remainingTasks := len(tasks) % workers

	for i := 0; i < workers; i++ {
		// divide tasks among workers
		start := i * tasksPerWorker
		end := start + tasksPerWorker

		// dispatch tasks to workers
		go func(index int, tasks []string) {
			for _, task := range tasks[index:end] {
				worker(task)
			}
		}(i, tasks)
	}

	// block main thread until all workers are finished
	// Similarly, this can be replaced with loops to wait for the completion of goroutines for a non-blocking solution
	time.Sleep(time.Duration(len(tasks)) * time.Second)
}

// goroutine and waitgroup optimization
func doLogisticsWithGoroutinesAndWaitGroup(workers int, tasks []string) {

	// create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// define the worker function
	worker := func(wg *sync.WaitGroup, task string) {
		defer wg.Done() // Mark this goroutine as complete
		fmt.Println("Worker processing:", task)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker completed:", task)
	}

	// Calculate the task workload per worker
	tasksPerWorker := len(tasks) / workers
	remainingTasks := len(tasks) % workers

	for i := 0; i < workers; i++ {
		// divide tasks among workers
		start := i * tasksPerWorker
		end := start + tasksPerWorker

		if i == workers-1 {
			end += remainingTasks
		}

		// launch a goroutine for each task
		for _, task := range tasks[start:end] {
			wg.Add(1) // Add 1 to the WaitGroup for each goroutine
			go worker(&wg, task)
		}
	}

	wg.Wait() // Block the main thread until all workers are complete
}
