package main

import (
	"fmt"
	"sync"
	"time"
)

// SharedCounter struct using RWMutex
type SharedCounter struct {
	count int
	rwmu  sync.RWMutex
}

// Increment method for incrementing the counter (writer operation)
func (sc *SharedCounter) Increment() {
	sc.rwmu.Lock()
	sc.count++
	sc.rwmu.Unlock()
}

// Decrement method for decrementing the counter (writer operation)
func (sc *SharedCounter) Decrement() {
	sc.rwmu.Lock()
	sc.count--
	sc.rwmu.Unlock()
}

// ReadCounter method for reading the counter (reader operation)
func (sc *SharedCounter) ReadCounter() int {
	sc.rwmu.RLock()
	defer sc.rwmu.RUnlock()
	return sc.count
}

func main() {
	var sc SharedCounter

	// Number of concurrent reader goroutines
	const numReaders = 100

	// Number of concurrent writer goroutines
	const numWriters = 5

	wg := &sync.WaitGroup{}
	wg.Add(numReaders + numWriters)

	//启动读取器
	for i := 0; i < numReaders; i++ {
		go func() {
			for j := 0; j < 100000; j++ { // Each reader thread reads 100000 times
				currentCount := sc.ReadCounter()
				time.Sleep(time.Microsecond * 50) //introduce some sleep to simulate various read speeds
			}
			wg.Done()
		}()
	}

	//启动写入器
	for i := 0; i < numWriters; i++ {
		go func() {
			for j := 0; j < 1000; j++ { // Each writer thread writes 1000 times
				sc.Increment()
				time.Sleep(time.Microsecond * 100) // introduce some sleep to simulate various write speeds
				sc.Decrement()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	// Final result
	fmt.Println("Final Counter Value:", sc.ReadCounter())
}
