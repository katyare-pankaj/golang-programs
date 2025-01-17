package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	day := t.Day()
	month := t.Month()
	year := t.Year()

	frenchMonths := []string{"", "Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"}
	frenchDays := []string{"", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"}

	fmt.Printf("Today is: %s, %d %s %d\n", frenchDays[int(t.Weekday())+1], day, frenchMonths[int(month)], year)
}
