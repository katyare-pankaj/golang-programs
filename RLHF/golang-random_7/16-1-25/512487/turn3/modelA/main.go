package main

import (
	"fmt"
	"time"
)

// Predefined arrays for French day and month names
var frenchDays = []string{
	"", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche",
}
var frenchMonths = []string{
	"", "Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
	"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
}

func frenchDateFormat(t time.Time) string {
	// Get the day of the week, month, and year from the time.Time object
	dayOfWeek := int(t.Weekday()) + 1 // Adjust for 0-based index
	month := t.Month()
	year := t.Year()

	// Format the date using French day and month names
	formattedDate := fmt.Sprintf("%s %d, %s %d", frenchDays[dayOfWeek], t.Day(), frenchMonths[month], year)
	return formattedDate
}

func main() {
	// Get the current time
	now := time.Now()

	// Format and print the current date in French
	frenchDate := frenchDateFormat(now)
	fmt.Println("Aujourd'hui est :", frenchDate)
}
