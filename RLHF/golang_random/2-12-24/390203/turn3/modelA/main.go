package main

import (
	"fmt"
	"sync"
)

var counter int
var mtx sync.Mutex

func incrementCounter() {
	mtx.Lock()
	counter++
	mtx.Unlock()
}

func main() {
	var wg sync.WaitGroup
	const numGoroutines = 1000
	const iterations = 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < iterations; j++ {
				incrementCounter()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter) // Should print 1000000
}
