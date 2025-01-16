package main

import (
	"fmt"
	"time"
)

var monthNamesFrench = [...]string{
	"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
	"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
}

func formatDateFrench(t time.Time) string {
	day := t.Day()
	month := monthNamesFrench[t.Month()-1]
	year := t.Year()
	return fmt.Sprintf("%02d %s %04d", day, month, year)
}

func main() {
	now := time.Now()
	fmt.Println(formatDateFrench(now))
}
