package main

import (
	"bufio"
	"fmt"
	"os"
)

// readFile reads the contents of a file and returns its lines as a slice of strings.
func readFile(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // Return the error if the file cannot be opened
	}
	defer file.Close() // Ensure the file is closed when done

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Append each line to the slice
	}

	if err := scanner.Err(); err != nil {
		return nil, err // Return any scanning errors
	}

	return lines, nil // Return the lines read from the file
}

// readFileData reads data from a specified file and prints an error message if it fails.
func readFileData(filename string, fileErrorMsg string) ([]string, error) {
	data, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading", fileErrorMsg, "file:", err)
		return nil, err
	}
	return data, nil
}

func main() {
	empFileData, err := readFileData("employees.txt", "employee")
	if err != nil {
		return // Exit if there's an error
	}

	roleFileData, err := readFileData("roles.txt", "role")
	if err != nil {
		return // Exit if there's an error
	}

	// Print the data read from both files
	fmt.Println("Employees:", empFileData)
	fmt.Println("Roles:", roleFileData)
}
