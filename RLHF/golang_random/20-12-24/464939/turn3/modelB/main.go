package main

import (
	"fmt"
)

func main() {
	studentGrades := map[string]map[string]int{
		"Alice":   {"Math": 80, "Science": 90, "History": 75},
		"Bob":     {"Math": 95, "Science": 88, "English": 92},
		"Charlie": {"Math": 78, "Science": 92, "Geography": 85},
	}

	for studentName, subjectGrades := range studentGrades {
		fmt.Printf("\nStudent Name: %s\n", studentName)
		for subject, grade := range subjectGrades {
			fmt.Printf("  %s: %d\n", subject, grade)
		}
	}
}
