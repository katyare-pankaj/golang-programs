package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const datasetSize = 10000000 // Adjust this size for testing

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
		return
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Generate a large dataset
	data := generateData(datasetSize)

	// Process the data using range loops
	sum := processData(data)

	fmt.Printf("Processed sum: %d\n", sum)
}

// generateData creates a slice with random numbers
func generateData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	return data
}

// processData processes the dataset and returns a computed sum
func processData(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}
