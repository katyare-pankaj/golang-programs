package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second) // Simulate a long operation
		fmt.Println("Operation completed.")
	}()

	time.Sleep(2 * time.Second) // Wait for 2 seconds
	if wg.Wait(0) == sync.ErrWaitTimeout {
		fmt.Println("Timed out waiting for operation.")
	}
}
