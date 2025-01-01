package main

import (
	"fmt"
)

// Student struct
type Student struct {
	Name  string
	Grade float64
}

func main() {
	// Initialize a slice of Student structs
	students := []Student{
		{Name: "Alice", Grade: 88.5},
		{Name: "Bob", Grade: 76.2},
		{Name: "Charlie", Grade: 92.0},
		{Name: "David", Grade: 85.8},
		{Name: "Eve", Grade: 67.9},
	}

	minGrade := 80.0 // Define the minimum grade threshold

	// Filter the students slice to include only students with grade greater than or equal to minGrade
	filteredStudents := filterStudents(students, func(s Student) bool {
		return s.Grade >= minGrade
	})

	// Print the filtered students
	fmt.Println("Filtered Students:")
	for _, student := range filteredStudents {
		fmt.Printf("%s: %.2f\n", student.Name, student.Grade)
	}
}

// filterStudents function takes a slice of Student structs and a predicate function as input
// and returns a new slice containing only those structs for which the predicate function returns true.
func filterStudents(students []Student, predicate func(Student) bool) []Student {
	var filtered []Student
	for _, student := range students {
		if predicate(student) {
			filtered = append(filtered, student)
		}
	}
	return filtered
}
