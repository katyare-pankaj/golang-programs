package main

import (
	"fmt"
)

func readFile(filename string) ([]string, error) {
	return nil, fmt.Errorf("read file error") // Placeholder for actual file reading error
}

func readFileData(filename string, fileErrorMsg string) ([]string, error) {
	data, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading ", fileErrorMsg, "file:", err)
		return nil, err
	}
	return data, nil
}

func main() {
	empFileData, err := readFileData("employees.txt", "employee")
	if err != nil {
		return
	}

	roleFileData, err := readFileData("roles.txt", "role")
	if err != nil {
		return
	}
}
