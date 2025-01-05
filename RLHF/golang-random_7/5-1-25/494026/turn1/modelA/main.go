package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// Simulate some work with time.Sleep
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Duration(id*2) * time.Second) // Simulate random work time
	fmt.Printf("Worker %d finished\n", id)
	wg.Done()
}

func main() {
	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 5

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers have finished.")

	// Start CPU and memory profiling
	cpuProfileFile := "cpu_profile.prof"
	memProfileFile := "mem_profile.prof"

	// Start CPU profiling
	f, err := os.Create(cpuProfileFile)
	if err != nil {
		log.Fatalf("could not create CPU profile: %v", err)
	}
	defer f.Close()
	if err := runtime.CPUProfile(f); err != nil {
		log.Fatalf("could not start CPU profile: %v", err)
	}
	defer runtime.CPUProfile(nil)

	// Start memory profiling
	f, err = os.Create(memProfileFile)
	if err != nil {
		log.Fatalf("could not create memory profile: %v", err)
	}
	defer f.Close()
	runtime.MemProfile(f)

	// Analyze the profiles using 'go tool pprof'
	// Example: go tool pprof -svg cpu_profile.prof > cpu_profile.svg
	// Example: go tool pprof -svg mem_profile.prof > mem_profile.svg
	fmt.Printf("CPU profile written to %s\n", cpuProfileFile)
	fmt.Printf("Memory profile written to %s\n", memProfileFile)
}
