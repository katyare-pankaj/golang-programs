package main

import (
	"fmt"
	"time"
)

const frenchDateFormat = "Lundi, 2 Janvier 2006"

var frenchDays = [7]string{
	"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche",
}
var frenchMonths = [12]string{
	"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
	"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
}

func frenchDateFormat(t time.Time) string {
	// Use the French date format string and the French day and month names
	return t.Format(frenchDateFormat)
}

func main() {
	now := time.Now()
	frenchDate := frenchDateFormat(now)
	fmt.Println("Today's date in French:", frenchDate)
}
