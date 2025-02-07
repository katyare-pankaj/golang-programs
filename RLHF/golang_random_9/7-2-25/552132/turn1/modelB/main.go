package main

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	slice []int
	mu    sync.Mutex
}

func (s *SafeSlice) Append(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, value)
}
func (s *SafeSlice) Get() []int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]int{}, s.slice...) // Return a copy
}
func worker(wg *sync.WaitGroup, safeSlice *SafeSlice, id int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		safeSlice.Append(id*100 + i)
	}
}
func main() {
	const numWorkers = 5
	var wg sync.WaitGroup
	safeSlice := &SafeSlice{}

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, safeSlice, i)
	}

	wg.Wait()
	result := safeSlice.Get()
	fmt.Println("Final result:", result)
}
