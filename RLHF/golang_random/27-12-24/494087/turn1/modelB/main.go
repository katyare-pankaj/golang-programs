package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// define a workItem struct
type workItem struct {
	id   int
	data []byte
}

// processWorkItem function
func processWorkItem(wg *sync.WaitGroup, mux *sync.Mutex, ch chan workItem, item workItem) {
	defer wg.Done()
	// simulate work using time.Sleep
	time.Sleep(time.Duration(100*item.id) * time.Millisecond)

	// lock the mutex for writing
	mux.Lock()
	// append the item to the channel
	ch <- item
	// unlock the mutex
	mux.Unlock()
}

func main() {
	// set the number of goroutines to use
	numGoroutines := 10
	// set the number of work items to process
	numWorkItems := 1000

	// allocate memory for the work items
	workItems := make([]workItem, numWorkItems)
	for i := 0; i < numWorkItems; i++ {
		workItems[i] = workItem{id: i, data: make([]byte, 1024)}
	}

	// create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	// create a Mutex to protect shared resources
	var mux sync.Mutex
	// create a channel to pass work items
	ch := make(chan workItem)

	// measure memory usage before starting the work
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// start the goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for item := range ch {
				// process the work item
				processWorkItem(&wg, &mux, ch, item)
			}
		}()
	}

	// send the work items to the goroutines
	for _, item := range workItems {
		ch <- item
	}

	// close the channel to signal the end of work
	close(ch)

	// wait for all goroutines to finish
	wg.Wait()

	// measure memory usage after the work is done
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	// calculate the memory usage difference
	memoryUsage := memStatsAfter.Alloc - memStatsBefore.Alloc

	fmt.Printf("Memory usage: %d bytes\n", memoryUsage)
}
