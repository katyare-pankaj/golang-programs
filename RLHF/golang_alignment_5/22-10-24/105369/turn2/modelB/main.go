package main

import (
	"fmt"
	"strings"
)

// measureIndentation calculates the indentation level of each line in the given code.
func measureIndentation(code string) []int {
	indentationLevels := []int{}
	lines := strings.Split(code, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimLeft(line, " \t")
		indentationLevel := len(line) - len(trimmedLine)
		indentationLevels = append(indentationLevels, indentationLevel)
	}
	return indentationLevels
}

// visualizeIndentation uses indentation levels to print the code with spaces or tabs to represent the hierarchy.
func visualizeIndentation(code string, indentationLevels []int) {
	lines := strings.Split(code, "\n")
	for i, line := range lines {
		// Generate the string of spaces based on the indentation level
		indentation := strings.Repeat(" ", indentationLevels[i])
		fmt.Printf("%s%s\n", indentation, line)
	}
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

	fmt.Println("Visualized Code with Indentation:")
	visualizeIndentation(codeExample, indentationLevels)
}
