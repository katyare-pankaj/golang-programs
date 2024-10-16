package main

import (
	"fmt"
	"time"
)

// Define a base interface for DataProcessor
type DataProcessor interface {
	Process() string
}

// Implement a concrete type for TextDataProcessor
type TextDataProcessor struct {
	data string
}

func (t TextDataProcessor) Process() string {
	// Perform text processing here
	return fmt.Sprintf("Processed Text: %s", t.data)
}

// Implement another concrete type for NumberDataProcessor
type NumberDataProcessor struct {
	data int
}

func (n NumberDataProcessor) Process() string {
	// Perform number processing here
	return fmt.Sprintf("Processed Number: %d", n.data)
}

func processData(dp DataProcessor) string {
	return dp.Process()
}

func main() {
	// Before using polymorphism
	textData := "Hello RPA!"
	numberData := 100

	startTime := time.Now()

	processedText := processTextData(textData)
	processedNumber := processNumberData(numberData)

	fmt.Println("Processed Text:", processedText)
	fmt.Println("Processed Number:", processedNumber)

	elapsedTime := time.Since(startTime)
	fmt.Println("Execution time before polymorphism:", elapsedTime)

	// After using polymorphism
	startTime = time.Now()

	textProcessor := TextDataProcessor{data: textData}
	numberProcessor := NumberDataProcessor{data: numberData}

	processedText = processData(textProcessor)
	processedNumber = processData(numberProcessor)

	fmt.Println("Processed Text:", processedText)
	fmt.Println("Processed Number:", processedNumber)

	elapsedTime = time.Since(startTime)
	fmt.Println("Execution time after polymorphism:", elapsedTime)
}
