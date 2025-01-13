package main

import (
	"fmt"
	"sync"
)

// Function that simulates interacting with a travel API
func travelAPI(apiName string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Calling API: %s\n", apiName)
	// Simulate some work being done
	// ...
}

func main() {
	var wg sync.WaitGroup
	apis := []string{"API1", "API2", "API3", "API4"}

	wg.Add(len(apis))
	for _, api := range apis {
		go travelAPI(api, &wg)
	}

	wg.Wait()
	fmt.Println("All APIs have been called.")
}
