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
	taskHour, taskMinute := 10, 0
	taskTime := time.Date(now.Year(), now.Month(), now.Day(), taskHour, taskMinute, 0, 0, loc)

	// If the task time has already passed, schedule it for the next day
	if taskTime.Before(now) {
		taskTime = taskTime.AddDate(0, 0, 1) // Add one day
	}

	// Calculate the time duration until the task should run
	duration := taskTime.Sub(now)

	// Schedule the task to run at the calculated time
	time.AfterFunc(duration, func() {
		fmt.Println("Running task in New York time:", time.Now().In(loc))
	})

	select {} // Infinite loop to keep the program running
}
