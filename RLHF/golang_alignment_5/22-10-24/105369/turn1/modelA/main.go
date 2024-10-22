package main

import (
	"fmt"
	"strings"
)

func measureIndentation(code string) []int {
	indentationLevels := []int{}
	lines := strings.Split(code, "\n")

	for _, line := range lines {
		// Trim leading spaces and tabs
		trimmedLine := strings.TrimLeft(line, " \t")

		// Calculate the indentation level by counting the number of leading spaces
		indentationLevel := len(line) - len(trimmedLine)
		indentationLevels = append(indentationLevels, indentationLevel)
	}
	return indentationLevels
}

func main() {
	codeExample := `
		package main
		import ("fmt")

		func main() {
			fmt.Println("Hello, World!")
			if true {
				fmt.Println("Indented block")
			}
		}
	`
	indentationLevels := measureIndentation(codeExample)

	fmt.Println("Indentation Levels:")
	for i, level := range indentationLevels {
		fmt.Printf("%d: %d\n", i+1, level)
	}
}
