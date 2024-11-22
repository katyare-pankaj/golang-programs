package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func work(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "cancelled")
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("worker", id, "working")
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(ctx, &wg, i)
	}
	// Simulate some work that might take a long time
	time.Sleep(5 * time.Second)
	// Cancel the context after 5 seconds
	cancel()
	wg.Wait()
	fmt.Println("all workers cancelled")
}
