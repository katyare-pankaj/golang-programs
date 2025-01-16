package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<file_path>")
		os.Exit(1)
	}

	// Open the file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// Defer the closure of the file to ensure it gets closed after main() returns
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Print the file content
	fmt.Println(string(content))
}
