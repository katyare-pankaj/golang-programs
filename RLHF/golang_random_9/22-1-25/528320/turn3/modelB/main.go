package main

import (
	"fmt"
	"go-programs/RLHF/golang_random_9/22-1-25/528320/turn2/modelA/funclib"
)

type Employee struct {
	Name   string
	Salary int
}

func main() {
	employees := []Employee{
		{Name: "Alice", Salary: 50000},
		{Name: "Bob", Salary: 45000},
		{Name: "Charlie", Salary: 60000},
		{Name: "David", Salary: 40000},
		{Name: "Eve", Salary: 55000},
	}

	//Calculating the average salary
	averageSalary := funclib.Reduce(funclib.Map(employees, func(emp Employee) int {
		return emp.Salary
	}),
		func(acc, val int) int {
			return acc + val
		}, 0) / len(employees)

	fmt.Println("Average Salary:", averageSalary)

	//Filtering employees with salary greater than or equal to the average salary
	employeesWithHighSalary := funclib.Filter(employees, func(emp Employee) bool {
		return emp.Salary >= averageSalary
	})

	fmt.Println("Employees with salary greater than or equal to average:")
	for _, emp := range employeesWithHighSalary {
		fmt.Println(emp.Name, ":", emp.Salary)
	}
}
