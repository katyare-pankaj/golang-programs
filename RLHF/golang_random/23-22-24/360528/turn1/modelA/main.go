package main

import (
	"fmt"
	"time"
)

// task is a function that simulates performing some task.
func task() {
	fmt.Println("Task is running.")
}

// main demonstrates scheduling a task at 10 AM in "America/New_York".
func main() {
	// Define the time zone location
	newYorkLocation, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading time zone:", err)
		return
	}

	// Calculate the next occurrence of 10 AM EST/EDT
	now := time.Now().In(newYorkLocation)
	next10AM := now.AddDate(0, 0, 1-now.Hour()%24).Add(10 * time.Hour)

	// Create a ticker that fires every 24 hours at 10 AM EST/EDT
	ticker := time.NewTicker(time.Until(next10AM))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			task()

			// Print the current time in "America/New_York"
			fmt.Println("Current time in New York:", time.Now().In(newYorkLocation).Format("2006-01-02 15:04:05 -0700"))
		}
	}
}
