package main

import (
	"fmt"
	"sync"
	"time"
)

func processString(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing string:", s)
	time.Sleep(time.Duration(1) * time.Second) // Simulate work
}

func main() {
	stringsToProcess := []string{"Apple", "Banana", "Cherry", "Orange", "Pineapple"}

	var wg sync.WaitGroup
	wg.Add(len(stringsToProcess))

	for _, s := range stringsToProcess {
		// Capture the current value of s using a closure
		go func(str string) {
			processString(str, &wg)
		}(s)
	}

	fmt.Println("Main function waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("All strings processed.")
}
