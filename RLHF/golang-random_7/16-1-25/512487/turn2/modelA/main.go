package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date
	now := time.Now()

	// Format the date in a way that can be easily understood and adapted for French
	dateStr := fmt.Sprintf("%2d %v %d", now.Day(), now.Month(), now.Year())

	// Map the months to their French names
	months := []string{"", "janvier", "février", "mars", "avril", "mai", "juin",
		"juillet", "août", "septembre", "octobre", "novembre", "décembre"}

	// Replace the English month name with its French equivalent
	dateStr = fmt.Sprintf("%s %s %d", dateStr[:5], months[now.Month()], now.Year())

	// Ensure proper punctuation
	dateStr = dateStr[:len(dateStr)-1] + "."

	fmt.Println("Date in French:", dateStr)
}
