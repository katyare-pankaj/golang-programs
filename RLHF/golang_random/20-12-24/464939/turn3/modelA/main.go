package main

import "fmt"

func main() {
	// Define the nested map where the outer map's keys are student names
	// and the inner maps hold grades for each subject.
	studentGrades := map[string]map[string]int{
		"Alice":   {"Math": 85, "Science": 90, "History": 88},
		"Bob":     {"Math": 78, "Science": 85, "History": 80},
		"Charlie": {"Math": 92, "Science": 95, "History": 90},
	}

	// Iterate over each student in the outer map.
	for studentName, grades := range studentGrades {
		// Print the student's name.
		fmt.Printf("Student: %s\n", studentName)

		// Iterate over each subject and grade in the inner map.
		for subject, grade := range grades {
			// Print the subject and grade.
			fmt.Printf("  %s: %d\n", subject, grade)
		}
		// Optionally, add a blank line for readability.
		fmt.Println()
	}
}
