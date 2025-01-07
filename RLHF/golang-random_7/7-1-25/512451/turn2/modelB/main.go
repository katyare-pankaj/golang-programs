package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numGenerators = 5
	logBufferSize = 100
	logInterval   = 100 * time.Millisecond
)

func generateLogs(wg *sync.WaitGroup, logCh chan string) {
	defer wg.Done()
	for {
		log := fmt.Sprintf("Log message generated at %v", time.Now())
		logCh <- log
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
func processLogs(wg *sync.WaitGroup, logCh chan string) {
	defer wg.Done()
	for log := range logCh {
		fmt.Println(log)
		// Simulate processing time
		time.Sleep(logInterval)
	}
}
func main() {
	var wg sync.WaitGroup
	logCh := make(chan string, logBufferSize)
	wg.Add(numGenerators + 1)
	// Start generating log messages
	for i := 0; i < numGenerators; i++ {
		go generateLogs(&wg, logCh)
	}
	// Start processing log messages
	go processLogs(&wg, logCh)
	wg.Wait()
}
