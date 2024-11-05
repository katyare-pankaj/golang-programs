package main

import (
	"bufio"
	"fmt"
	"os"
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
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // Return the error if file cannot be opened
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Append each line to the slice
	}

	if err := scanner.Err(); err != nil {
		return nil, err // Return any scanning errors
	}

	return lines, nil // Return the lines read from the file
}

func main() {
	empFile := readEmployeeFile("employees.txt")
	roleFile := readRoleFile("roles.txt")

	fmt.Println("Employees:", empFile)
	fmt.Println("Roles:", roleFile)
}
