package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

var sharedResource int = 0
var rwmutex sync.RWMutex

func write(ch <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		<-ch // Wait for a signal to perform a write
		rwmutex.Lock()
		sharedResource += 1
		rwmutex.Unlock()
		fmt.Println("Resource value updated:", sharedResource)
		time.Sleep(time.Duration(int64(rand.Intn(100))+1) * time.Millisecond)
	}
}

func read(ch <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		<-ch // Wait for a signal to perform a read
		rwmutex.RLock()
		fmt.Println("Resource value read:", sharedResource)
		rwmutex.RUnlock()
		time.Sleep(time.Duration(int64(rand.Intn(100))+1) * time.Millisecond)
	}
}

func main() {
	const (
		writers = 5
		readers = 10
	)

	var wg sync.WaitGroup

	ch := make(chan struct{}, writers+readers)

	// Start writer goroutines
	for i := 0; i < writers; i++ {
		wg.Add(1)
		go write(ch, &wg)
	}

	// Start reader goroutines
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go read(ch, &wg)
	}

	// Signal goroutines to perform operations
	for i := 0; i < writers+readers; i++ {
		ch <- struct{}{}
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
