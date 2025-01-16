package main

import (
	"fmt"
	"time"
)

// func frenchDateFormat(t time.Time) string
// Takes a `time.Time` object and returns a formatted date string in French locale.
func frenchDateFormat(t time.Time) string {
	// Define French day names
	frenchDays := [7]string{
		"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche",
	}
	// Define French month names
	frenchMonths := [12]string{
		"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
		"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
	}
	// Format using custom strings
	f := "%s, %d %s %d"
	return fmt.Sprintf(f, frenchDays[t.Weekday()], t.Day(), frenchMonths[t.Month()-1], t.Year())
}

func main() {
	// Example time
	now := time.Now()

	// Format the time to a French date
	frenchDate := frenchDateFormat(now)
	fmt.Println("Today's date in French:", frenchDate)
}
