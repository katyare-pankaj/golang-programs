package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go say("hello", &wg)
	go say("greetings", &wg)
	go say("bonjour", &wg)
	wg.Wait()
	fmt.Println("All goroutines are done.")
}
