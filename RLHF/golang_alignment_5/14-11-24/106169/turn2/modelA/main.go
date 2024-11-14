package main

import (
	"fmt"
	"io/ioutil"
)

func processFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Process the data here
	fmt.Println("Data processed:", string(data))
	return nil
}

func main() {
	filePath := "example.txt"
	err := processFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
