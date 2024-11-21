package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	task := func() error {
		// Simulate a task that might fail sometimes
		time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
		if rand.Intn(2) == 0 {
			return errors.New("task failed randomly")
		}
		fmt.Println("Task executed successfully.")
		return nil
	}

	errChan := make(chan error)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go scheduleTask(task, 3*time.Second, errChan)

	for {
		select {
		case err := <-errChan:
			fmt.Println("Received error:", err)
		case <-stopChan:
			fmt.Println("Stopping scheduler...")
			return
		}
	}
}

func scheduleTask(task func() error, interval time.Duration, errChan chan<- error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := task(); err != nil {
				errChan <- err
			}
		case <-time.After(interval * 2):
			errChan <- errors.New("task execution timed out")
		}
	}
}
