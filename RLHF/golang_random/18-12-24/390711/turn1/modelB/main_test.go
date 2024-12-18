package main

import (
	"sync"
	"testing"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
}

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg)
	wg.Wait()
}
