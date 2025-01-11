package main

import (
	"fmt"
	"sync"
)

var (
	sharedData = make([]int, 10)
	mutex1     sync.Mutex
	mutex2     sync.Mutex
	wg         sync.WaitGroup
)

func doWork(index int) {
	defer wg.Done()

	mutex1.Lock()
	defer mutex1.Unlock()

	mutex2.Lock()
	defer mutex2.Unlock()

	sharedData[index]++
}

func main() {

	for i := 0; i < len(sharedData); i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()

	fmt.Println("All work complete.")
	fmt.Println("Final shared data:", sharedData)
}
