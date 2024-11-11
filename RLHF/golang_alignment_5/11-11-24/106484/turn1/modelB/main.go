package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Code block with potential bottleneck
	sharedResource := 0

	go func() {
		incrementResource(&sharedResource, 10000)
		wg.Done()
	}()

	go func() {
		incrementResource(&sharedResource, 10000)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final shared resource value:", sharedResource)
}

func incrementResource(resource *int, iterations int) {
	for i := 0; i < iterations; i++ {
		// Using atomic operation to avoid race condition and ensure consistency
		*resource++
	}
}
