package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkers(t *testing.T) {
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers completed.")
}
