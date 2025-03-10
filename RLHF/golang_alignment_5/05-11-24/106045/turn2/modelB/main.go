package main

import (
	"fmt"
)

func readEmployeeFile(filename string) []string {
	employees, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading employee file:", err)
		return nil
	}
	return employees
}

func readRoleFile(filename string) []string {
	names, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading role file:", err)
		return nil
	}
	return names
}

func readFile(filename string) ([]string, error) {
	return nil, fmt.Errorf("read file error") // Placeholder for actual file reading error
}

func main() {
	empFile := readEmployeeFile("employees.txt")
	roleFile := readRoleFile("roles.txt")
}
