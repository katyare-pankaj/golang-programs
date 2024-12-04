package main

import (
	"fmt"
	"log"
	"os"
)

func divideByZero(x float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
			// Optionally, you could rethrow the panic here if necessary
		}
	}()
	if x == 0 {
		panic("division by zero is not allowed")
	}
	return 1.0 / x
}

func readFile(filename string) string {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in readFile: %v\n", r)
			os.Exit(1) // Exit program if file read fails
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}
	defer file.Close() // Ensure file is closed

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
	}
	return string(buf[:n])
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Unrecovered panic in main: %v\n", r)
		}
	}()

	result := divideByZero(0)
	fmt.Println("Result of division:", result) // Output will be "Recovered from panic: division by zero is not allowed"

	fileContent := readFile("example.txt")
	fmt.Println("File content:", fileContent) // Output will be "Recovered from panic in readFile: failed to open file: no such file or directory"
}
