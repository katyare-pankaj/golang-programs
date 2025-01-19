package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

const (
	mapSize  = 1_000_000
	dataSize = 10_000_000
)

// initializes the input map with data
func populateMap() map[int]int {
	data := make(map[int]int, mapSize)
	for i := 0; i < mapSize; i++ {
		data[i] = i
	}
	return data
}

// This function processes the data from the input map
func processData(input map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < dataSize; i++ {
		// perform some operation on the data
		_, _ = input[i%mapSize] // simulate a key lookup
	}
}

func main() {
	// Disable GC for a more precise profile
	runtime.GC()

	// Create the input map
	inputMap := populateMap()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Start CPU profile
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	go processData(inputMap, wg)
	wg.Wait()

	fmt.Println("Processing completed")
}
