package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processBatch(batch []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	fmt.Println("Processed batch:", batch)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const batchSize = 1000
	const numBatches = 10

	for i := 0; i < numBatches; i++ {
		batch := make([]int, batchSize)
		for j := 0; j < batchSize; j++ {
			batch[j] = rand.Intn(1000)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go processBatch(batch, &wg)
		wg.Wait()
	}
}
