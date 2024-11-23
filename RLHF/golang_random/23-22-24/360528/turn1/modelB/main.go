package main

import (
	"fmt"
	"time"
)

func main() {
	// Load the desired time zone location
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	// Get the current time in the specified time zone
	now := time.Now().In(loc)

	// Define the desired time for the task to run (e.g., 10 AM in New York)
	taskTime := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, loc)

	// If the task time has already passed, schedule it for the next day
	if taskTime.Before(now) {
		taskTime = taskTime.Add(24 * time.Hour)
	}

	// Create a timer to trigger the task at the desired time
	timer := time.NewTimer(taskTime.Sub(now))

	// Channel to receive the timer event
	taskChan := make(chan bool)

	// Go routine to handle the task
	go func() {
		<-timer.C
		fmt.Println("Running task in New York time:", time.Now().In(loc))
		taskChan <- true
	}()

	// Wait for the task to complete
	<-taskChan
}
