package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedResource struct {
	data map[string]string
	once sync.Once
}

func (r *SharedResource) Initialize() {
	r.once.Do(func() {
		fmt.Println("Initializing shared resource...")
		r.data = make(map[string]string)
		time.Sleep(2 * time.Second) // Simulate initialization time
		fmt.Println("Shared resource initialized.")
	})
}

func (r *SharedResource) GetData(key string) string {
	r.Initialize()
	return r.data[key]
}

func main() {
	var sharedResource SharedResource

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 1: Getting value for key 'a': %s\n", sharedResource.GetData("a"))
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 2: Getting value for key 'b': %s\n", sharedResource.GetData("b"))
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(10 * time.Second) // Wait for both goroutines to complete
}
