// main.go
package main

import (
	"fmt"
	"go-programs/RLHF/golang_random_9/22-1-25/528320/turn2/modelB/fplib"
)

// Employee represents an employee with a name and salary.
type Employee struct {
	Name   string
	Salary float64
}

func main() {
	// Sample data: List of employees with their salaries
	employees := []Employee{
		{"Alice", 60000},
		{"Bob", 55000},
		{"Charlie", 70000},
		{"David", 45000},
		{"Eva", 75000},
	}

	// Map: Extract the list of salaries
	salaries := fplib.Map(employees, func(e Employee) float64 { return e.Salary })
	fmt.Println("Salaries:", salaries)

	// Reduce: Calculate the total sum of all salaries
	totalSalaries := fplib.Reduce(salaries, func(acc, salary float64) float64 { return acc + salary }, 0.0)
	fmt.Printf("Total Salaries: %.2f\n", totalSalaries)

	// Filter: Find all employees with salaries above $60,000
	highEarners := fplib.Filter(employees, func(e Employee) bool { return e.Salary > 60000 })
	fmt.Println("Employees with salaries above $60,000:")
	for _, e := range highEarners {
		fmt.Println("-", e.Name)
	}

	// Reduce: Find the highest salary
	maxSalary := fplib.Reduce(salaries, func(max, salary float64) float64 {
		if salary > max {
			return salary
		}
		return max
	}, 0.0)
	fmt.Printf("Max Salary: %.2f\n", maxSalary)
}
