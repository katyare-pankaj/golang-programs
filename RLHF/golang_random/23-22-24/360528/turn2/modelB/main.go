package main

import (
	"fmt"
	"time"
)

func main() {
	// Load the desired time zone location.
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	// Get the current time in the specified time zone.
	now := time.Now().In(loc)

	// Define the desired time for the task to run (e.g., 6 AM in New York).
	taskTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, loc)

	// If the task time has already passed, schedule it for the next day.
	if taskTime.Before(now) {
		taskTime = taskTime.Add(24 * time.Hour)
	}

	fmt.Println("Task scheduled to run at:", taskTime)

	// Create a timer to trigger the task at the desired time.
	timer := time.NewTimer(taskTime.Sub(now))

	// Wait for the timer to trigger, and then run the task.
	<-timer.C
	fmt.Println("Running task...")
	time.Sleep(5 * time.Second) // Simulate task execution time
	fmt.Println("Task completed.")
}
