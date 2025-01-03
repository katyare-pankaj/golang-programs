package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, id int, delay time.Duration) {
	defer wg.Done()
	fmt.Printf("Task %d started with a delay of %s\n", id, delay)
	time.Sleep(delay)
	fmt.Printf("Task %d completed.\n", id)
}

func main() {
	var wg sync.WaitGroup
	numTasks := 3
	delays := make([]time.Duration, 0)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < numTasks; i++ {
		fmt.Printf("Enter the delay time for task %d (in seconds): ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		delayTime, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer value.")
			i--
			continue
		}

		delays = append(delays, time.Duration(delayTime)*time.Second)
	}

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go task(&wg, i+1, delays[i])
	}
	wg.Wait()
	fmt.Println("All tasks completed. Main function exiting...")
}
