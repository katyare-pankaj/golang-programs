package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func logEventsWithChannel() {
	logCh := make(chan string, 100)
	defer close(logCh)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range logCh {
			fmt.Println(msg)
		}
	}()

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		logCh <- fmt.Sprintf("Event %d", i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	logEventsWithChannel()
	wg.Wait()
}
