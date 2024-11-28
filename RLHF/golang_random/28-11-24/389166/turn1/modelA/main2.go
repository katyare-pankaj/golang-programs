package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

func logEventsWithMutex() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		mu.Lock()
		fmt.Println(fmt.Sprintf("Event %d", i))
		mu.Unlock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	logEventsWithMutex()
}
