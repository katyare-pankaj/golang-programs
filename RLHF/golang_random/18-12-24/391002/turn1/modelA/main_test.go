package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func processItem(wg *sync.WaitGroup, item string, result chan<- error) {
	defer wg.Done()

	if item == "" {
		result <- fmt.Errorf("item cannot be empty")
		return
	}

	// Simulate some work
	select {
	case <-time.After(time.Millisecond * 100):
	case <-time.After(time.Second): // Simulate timeout
		result <- fmt.Errorf("timeout processing item %s", item)
		return
	}

	result <- nil
}

func TestProcessItems(t *testing.T) {
	var wg sync.WaitGroup
	resultCh := make(chan error, 3)

	items := []string{"item1", "", "item3"}

	for _, item := range items {
		wg.Add(1)
		go processItem(&wg, item, resultCh)
	}

	wg.Wait()

	close(resultCh)

	for err := range resultCh {
		if err != nil {
			t.Errorf("processItem failed: %v", err)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	resultCh := make(chan error, 3)

	items := []string{"item1", "", "item3"}

	for _, item := range items {
		wg.Add(1)
		go processItem(&wg, item, resultCh)
	}

	wg.Wait()

	close(resultCh)

	for err := range resultCh {
		if err != nil {
			log.Fatalf("processItem failed: %v", err)
		}
	}
}
