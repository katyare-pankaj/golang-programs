package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.go") // Replace with your actual code file
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		indentationLevel := measureIndentationLevel(line)
		fmt.Printf("Indentation level: %d, Line: %s\n", indentationLevel, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}

// measureIndentationLevel measures the indentation level of a given line
// by counting the number of leading spaces or tabs.
func measureIndentationLevel(line string) int {
	indentationLevel := 0
	for _, char := range line {
		if char == ' ' || char == '\t' {
			indentationLevel++
		} else {
			break
		}
	}
	return indentationLevel
}
