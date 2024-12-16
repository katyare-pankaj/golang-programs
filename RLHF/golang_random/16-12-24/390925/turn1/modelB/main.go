package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var (
	wg       sync.WaitGroup
	dataSize = 1000
)

func matrixMultiply(m1, m2 [][]int) [][]int {
	result := make([][]int, dataSize, dataSize)
	for i := 0; i < dataSize; i++ {
		for j := 0; j < dataSize; j++ {
			for k := 0; k < dataSize; k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return result
}

func worker(id int) {
	defer wg.Done()
	// Initialize random matrices
	m1 := make([][]int, dataSize, dataSize)
	m2 := make([][]int, dataSize, dataSize)
	for i := 0; i < dataSize; i++ {
		for j := 0; j < dataSize; j++ {
			m1[i][j] = rand.Intn(100)
			m2[i][j] = rand.Intn(100)
		}
	}

	// Perform matrix multiplication
	_ = matrixMultiply(m1, m2)

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile:", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile:", err)
	}
	defer pprof.StopCPUProfile()

	numGoroutines := runtime.NumCPU()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()

	fmt.Println("All workers finished")
}
