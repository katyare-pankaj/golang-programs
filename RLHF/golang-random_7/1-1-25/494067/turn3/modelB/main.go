package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Grade float64
}

// This function takes a slice of Students and a grade threshold
// as input and returns a new slice of students with grades greater than or equal to the threshold.
func filterStudentsByGrade(students []Student, threshold float64) []Student {
	var filteredStudents []Student
	for _, student := range students {
		if student.Grade >= threshold {
			filteredStudents = append(filteredStudents, student)
		}
	}
	return filteredStudents
}

func main() {
	students := []Student{
		{Name: "Alice", Grade: 90.0},
		{Name: "Bob", Grade: 75.0},
		{Name: "Charlie", Grade: 88.0},
		{Name: "David", Grade: 66.0},
		{Name: "Eve", Grade: 92.0},
		{Name: "Frank", Grade: 78.0},
	}

	minGrade := 80.0 // specify the threshold grade

	filteredStudents := filterStudentsByGrade(students, minGrade)

	fmt.Println("Students with Grade >=", minGrade, ":")
	for _, student := range filteredStudents {
		fmt.Printf("%s: %.2f\n", student.Name, student.Grade)
	}
}
