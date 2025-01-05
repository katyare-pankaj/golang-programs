package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const (
	numTasks       = 10
	numbersPerTask = 1000000
)

var (
	totalSum int
)

func sumNumbers(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	var taskSum int
	for i := start; i < end; i++ {
		taskSum += rand.Intn(1000)
	}
	totalSum += taskSum
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create CPU profile file
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile file:", err)
		return
	}
	defer cpuFile.Close()
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("Error starting CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Create memory profile file
	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Error creating memory profile file:", err)
		return
	}
	defer memFile.Close()

	// Initialize waitgroup
	var wg sync.WaitGroup

	// Start concurrent tasks
	for i := 0; i < numTasks; i++ {
		start := i * numbersPerTask
		end := start + numbersPerTask
		wg.Add(1)
		go sumNumbers(start, end, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Write memory profile after all tasks are done
	pprof.WriteHeapProfile(memFile)

	// Print total sum
	fmt.Println("Total sum:", totalSum)

	// Analyze CPU profile using "go tool pprof cpu.prof"
	// Analyze memory profile using "go tool pprof mem.prof"
}
