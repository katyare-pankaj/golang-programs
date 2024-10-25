package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numIterations = 1000000000
	arraySize     = 1024
)

var (
	x [arraySize]float64
	y [arraySize]float64
	z [arraySize]float64
)

func init() {
	for i := 0; i < arraySize; i++ {
		x[i] = float64(i)
		y[i] = float64(i) * 2.0
	}
}

func matmulSIMD() {
	for i := 0; i < arraySize; i++ {
		for j := 0; j < arraySize; j++ {
			for k := 0; k < arraySize; k++ {
				z[i*arraySize+j] += x[i*arraySize+k] * y[k*arraySize+j]
			}
		}
	}
}

func matmulScalar() {
	for i := 0; i < arraySize; i++ {
		for j := 0; j < arraySize; j++ {
			temp := 0.0
			for k := 0; k < arraySize; k++ {
				temp += x[i*arraySize+k] * y[k*arraySize+j]
			}
			z[i*arraySize+j] = temp
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup

	wg.Add(2)

	start := time.Now()
	go func() {
		defer wg.Done()
		for i := 0; i < numIterations; i++ {
			matmulSIMD()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < numIterations; i++ {
			matmulScalar()
		}
	}()

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
}
