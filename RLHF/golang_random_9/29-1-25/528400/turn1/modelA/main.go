package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// Task simulates a work task by sleeping for a short time
func Task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement WaitGroup counter when the goroutine completes
	fmt.Printf("Task %d is starting...\n", id)
	time.Sleep(100 * time.Millisecond) // Simulate a task taking some time
	fmt.Printf("Task %d is completed.\n", id)
}

func main() {
	// Create a file to store the CPU profile
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup
	numTasks := 10

	// Set GOMAXPROCS to the number of available CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go Task(i, &wg)
	}

	// Start the pprof server for memory profiling
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wg.Wait() // Wait for all tasks to complete
	fmt.Println("All tasks completed.")
}
