package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const numThreads = 10
	const numIterations = 100000
	var sharedCounter int
	var lock sync.Mutex

	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		go func() {
			for j := 0; j < numIterations; j++ {
				lock.Lock()
				sharedCounter++
				lock.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final shared counter value:", sharedCounter)
}
