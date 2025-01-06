package main

import (
	"fmt"
	"regexp"
	"strings"
)

// StringTransformer encapsulates a sequence of transformation functions
type StringTransformer struct {
	Transformations []TransformFunction
}

// TransformFunction defines the signature for a transformation function
type TransformFunction func(string) string

// NewStringTransformer creates a new string transformer with optional transformations
func NewStringTransformer(transforms ...TransformFunction) *StringTransformer {
	return &StringTransformer{Transformations: transforms}
}

// Transform applies all transformations to the input string
func (st *StringTransformer) Transform(input string) string {
	output := input
	for _, transform := range st.Transformations {
		output = transform(output)
	}
	return output
}

// CapitalizeWord capitalizes the first letter of each word
func CapitalizeWord(input string) string {
	return strings.Title(input)
}

// ToLowercase converts the entire string to lowercase
func ToLowercase(input string) string {
	return strings.ToLower(input)
}

// RemoveNonAlphanumeric removes non-alphanumeric characters
func RemoveNonAlphanumeric(input string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	return re.ReplaceAllString(input, "")
}

// PadToLength pads the string with spaces to a specified length
func PadToLength(input string, length int) string {
	return fmt.Sprintf("%s%s", input, strings.Repeat(" ", length-len(input)))
}

func main() {
	// Create a pipeline with the specified transformations
	pipeline := NewStringTransformer(
		CapitalizeWord,        // Capitalize the first letter of each word
		RemoveNonAlphanumeric, // Remove non-alphanumeric characters
		ToLowercase,           // Convert to lowercase
		PadToLength,           // Pad to a specific length
	)

	// String to transform
	inputString := "transform This STRING! 123"

	// Transform the string using the pipeline
	outputString := pipeline.Transform(inputString)

	fmt.Println("Input:", inputString)
	fmt.Println("Output:", outputString)
}
