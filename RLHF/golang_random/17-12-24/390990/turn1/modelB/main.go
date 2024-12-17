package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	resource1 sync.Mutex
	resource2 sync.Mutex
)

func processResource1(ctx context.Context) {
	resource1.Lock()
	defer resource1.Unlock()

	fmt.Println("Processing Resource 1")

	// Simulate work
	time.Sleep(1 * time.Second)

	processResource2(ctx)
}

func processResource2(ctx context.Context) {
	resource2.Lock()
	defer resource2.Unlock()

	fmt.Println("Processing Resource 2")

	// Simulate work
	time.Sleep(1 * time.Second)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go processResource1(ctx)
	go processResource2(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Timeout reached. Exiting.")
	}
}
