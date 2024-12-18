package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go printHello("Alice", &wg)
	go printHello("Bob", &wg)

	wg.Wait()
	fmt.Println("All goroutines finished.")
}

func printHello(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("Hello from", name, "!")
		time.Sleep(time.Second * 1)
	}
}
