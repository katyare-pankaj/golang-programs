package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const numWorkers = 4
const sliceSize = 100000000

func main() {
	data := make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		data[i] = i + 1
	}

	var totalSum int64
	chunkSize := sliceSize / numWorkers

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		go worker(w, data, chunkSize, &totalSum, &wg)
	}

	wg.Wait()

	fmt.Println("Sum:", totalSum)
}

func worker(id int, data []int, chunkSize int, totalSum *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	start := id * chunkSize
	end := (id + 1) * chunkSize
	if end > len(data) {
		end = len(data)
	}

	var localSum int64
	for _, num := range data[start:end] {
		localSum += int64(num)
	}

	atomic.AddInt64(totalSum, localSum)
}
