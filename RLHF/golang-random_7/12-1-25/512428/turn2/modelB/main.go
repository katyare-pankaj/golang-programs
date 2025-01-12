package main

import (
	"fmt"
)

func generateFormattedStrings(format string, min int, max int) []string {
	var result []string
	for i := min; i <= max; i++ {
		str := fmt.Sprintf(format, i)
		result = append(result, str)
	}
	return result
}

func main() {
	pattern1 := "Day %02d"
	pattern2 := "Month %02d"
	pattern3 := "Year %04d"

	days := generateFormattedStrings(pattern1, 1, 31)
	months := generateFormattedStrings(pattern2, 1, 12)
	years := generateFormattedStrings(pattern3, 1900, 2023)

	fmt.Println("Days:")
	fmt.Println(days)

	fmt.Println("\nMonths:")
	fmt.Println(months)

	fmt.Println("\nYears:")
	fmt.Println(years)
}
