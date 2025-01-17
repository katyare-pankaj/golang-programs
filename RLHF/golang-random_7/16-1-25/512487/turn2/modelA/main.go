package main

import (
	"fmt"
	"time"
)

var frenchDays = []string{
	"", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche",
}
var frenchMonths = []string{
	"", "Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
	"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
}

func frenchDateFormat(t time.Time) string {
	day := frenchDays[int(t.Weekday())+1]
	month := frenchMonths[t.Month()]
	year := fmt.Sprintf("%04d", t.Year())
	return fmt.Sprintf("%s, %d %s %s", day, t.Day(), month, year)
}

func main() {
	// Example time
	now := time.Now()

	// Format the time to a French date
	frenchDate := frenchDateFormat(now)
	fmt.Println("Today's date in French:", frenchDate)
}
