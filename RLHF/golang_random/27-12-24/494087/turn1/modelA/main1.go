package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func waitGroupWithMutex(waitGroup *sync.WaitGroup, mutex *sync.Mutex, num int) {
	defer waitGroup.Done()
	for i := 0; i < num; i++ {
		mutex.Lock()
		time.Sleep(time.Nanosecond)
		mutex.Unlock()
	}
}

func waitGroupWithChannel(waitGroup *sync.WaitGroup, c <-chan interface{}) {
	defer waitGroup.Done()
	for range c {
		time.Sleep(time.Nanosecond)
	}
}

func main() {
	numGoroutines := 100
	numIterations := 100

	startMem := runtime.MemStats{}
	runtime.ReadMemStats(&startMem)

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	// Using WaitGroup and Mutex
	fmt.Println("Starting with Mutex...")
	for i := 0; i < numGoroutines; i++ {
		waitGroup.Add(1)
		go waitGroupWithMutex(&waitGroup, &mutex, numIterations)
	}
	waitGroup.Wait() // Wait for all goroutines to finish
	fmt.Println("Completed with Mutex.")

	midMem := runtime.MemStats{}
	runtime.ReadMemStats(&midMem)
	printMemStats(startMem, midMem, "Mutex")

	// Using WaitGroup and Channel
	fmt.Println("Starting with Channel...")
	c2 := make(chan interface{}, numIterations)
	for i := 0; i < numIterations; i++ {
		c2 <- struct{}{}
	}
	close(c2) // Close the channel so receivers stop waiting

	for i := 0; i < numGoroutines; i++ {
		waitGroup.Add(1)
		go waitGroupWithChannel(&waitGroup, c2)
	}
	waitGroup.Wait() // Wait for all goroutines to finish
	fmt.Println("Completed with Channel.")

	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	printMemStats(startMem, endMem, "Channel")
}

func printMemStats(start, end runtime.MemStats, name string) {
	alloc := end.Alloc - start.Alloc
	sys := end.Sys - start.Sys
	stackInuse := end.StackInuse - start.StackInuse
	heapAlloc := end.HeapAlloc - start.HeapAlloc
	heapSys := end.HeapSys - start.HeapSys

	fmt.Printf("Memory usage with %s:\n", name)
	fmt.Printf("Alloc: %d\n", alloc)
	fmt.Printf("Sys: %d\n", sys)
	fmt.Printf("StackInuse: %d\n", stackInuse)
	fmt.Printf("Heap Alloc: %d\n", heapAlloc)
	fmt.Printf("Heap Sys: %d\n", heapSys)
	fmt.Printf("====================\n")
}
