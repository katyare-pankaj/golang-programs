package main

import (
	"fmt"
	"sync"
)

var (
	// A fictitious, insecure counting function from a third-party library
	cnt     = 0
	cntLock = &sync.Mutex{}
)

func fakeLibCounter() {
	cntLock.Lock()
	cnt++
	cntLock.Unlock()
}

func main() {
	numGoroutines := 1000
	done := make(chan bool, numGoroutines)

	// launch goroutines to increment the fake counter
	for i := 0; i < numGoroutines; i++ {
		go func() {
			fakeLibCounter()
			done <- true
		}()
	}

	// wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	// output final count
	fmt.Printf("Final counter value: %d\n", cnt)
}
