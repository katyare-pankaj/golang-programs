package main

import (
	"runtime"
	"sync"
	"time"
)

func waitGroupWithMutex(waitGroup *sync.WaitGroup, c <-chan interface{}, num int) {
	defer waitGroup.Done()
	for i := 0; i < num; i++ {
		mutex.Lock()
		defer mutex.Unlock()
		time.Sleep(time.Nanosecond)
	}
}

func waitGroupWithChannel(waitGroup *sync.WaitGroup, c <-chan interface{}, num int) {
	defer waitGroup.Done()
	for i := 0; i < num; i++ {
		select {
		case <-c:
			time.Sleep(time.Nanosecond)
		}
	}
}

func main() {
	numGoroutines := 100000
	numIterations := 1000000

	startMem := runtime.MemStats{}
	runtime.ReadMemStats(&startMem)

	var waitGroup = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}

	// Using WaitGroup and Mutex
	c := make(chan interface{})
	for i := 0; i < numGoroutines; i++ {
		waitGroup.Add(1)
		go waitGroupWithMutex(waitGroup, c, numIterations)
		c <- "" // Start the goroutine
	}
	close(c)
	waitGroup.Wait()

	midMem := runtime.MemStats{}
	runtime.ReadMemStats(&midMem)
	printMemStats(startMem, midMem, "Mutex")

	// Using WaitGroup and Channel
	c2 := make(chan interface{}, numIterations*numGoroutines)
	for i := 0; i < numIterations*numGoroutines; i++ {
		c2 <- ""
	}

	for i := 0; i < numGoroutines; i++ {
		waitGroup.Add(1)
		go waitGroupWithChannel(waitGroup, c2, numIterations)
	}
	waitGroup.Wait()

	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	printMemStats(startMem, endMem, "Channel")
}

func printMemStats(start, end runtime.MemStats, name string) {
	alloc := end.Alloc - start.Alloc
	sys := end.Sys - start.Sys
	stack := end.Stack - start.Stack
	heapAlloc := end.HeapAlloc - start.HeapAlloc
	heapSys := end.HeapSys - start.HeapSys

	printf("Memory usage with %s:\n", name)
	printf("Alloc: %d\n", alloc)
	printf("Sys: %d\n", sys)
	printf("Stack: %d\n", stack)
	printf("Heap Alloc: %d\n", heapAlloc)
	printf("Heap Sys: %d\n", heapSys)
	printf("====================\n")
}
