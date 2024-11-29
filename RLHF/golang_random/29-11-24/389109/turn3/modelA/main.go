package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SharedCounter struct {
	counter int
	rwmu    sync.RWMutex
}

func (s *SharedCounter) Read() int {
	s.rwmu.RLock()
	defer s.rwmu.RUnlock()
	return s.counter
}

func (s *SharedCounter) Update(amount int) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()
	s.counter += amount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	counter := &SharedCounter{counter: 0}
	numGoroutines := 100
	const numReads = 100000

	// Helper function for read operations
	read := func(s *SharedCounter, id int) {
		for i := 0; i < numReads; i++ {
			value := s.Read()
			if i%1000 == 0 {
				go fmt.Printf("Goroutine %d: Read %d at iteration %d\n", id, value, i)
			}
		}
	}

	// Helper function for update operations
	update := func(s *SharedCounter) {
		for i := 0; i < numReads; i++ {
			amount := rand.Intn(10) - 5 // Random update amount between -5 and 5
			s.Update(amount)
			if i%1000 == 0 {
				go fmt.Printf("Goroutine %d: Updated counter to %d at iteration %d\n", rand.Intn(numGoroutines), s.Read(), i)
			}
		}
	}

	// Start reader goroutines
	for i := 0; i < numGoroutines; i++ {
		go read(counter, i)
	}

	// Start update goroutines
	for i := 0; i < numGoroutines/2; i++ {
		go update(counter)
	}

	// Allow goroutines to complete
	time.Sleep(3 * time.Second)

	fmt.Println("Final counter value:", counter.Read())
}
