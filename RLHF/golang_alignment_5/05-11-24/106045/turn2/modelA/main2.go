package main

import "fmt"

type Employee struct {
	Name   string
	Skills []string
}

func (e Employee) printCompetencies() {
	fmt.Println("Competencies:")
	for _, skill := range e.Skills {
		fmt.Println(skill)
	}
}

type Manager struct {
	Employee
	ReportingTeam []string
}

func main() {
	emp := Employee{
		Name:   "Alice",
		Skills: []string{"Leadership", "Management", "Communication"},
	}

	manager := Manager{
		Employee:      emp,
		ReportingTeam: []string{"Bob", "Charlie"},
	}

	emp.printCompetencies()
}
