package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

// Asynchronous file reader using CPS
func readFileAsyncCPS(filename string, callback func(error, []string)) {
	go func() {
		defer runtime.Gosched() // Yield to other goroutines
		var lines []string
		file, err := os.Open(filename)
		if err != nil {
			callback(err, nil)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			callback(err, nil)
			return
		}
		callback(nil, lines)
	}()
}

// Asynchronous file reader using async/await (Go 1.15+)
func readFileAsyncAwait(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func processLines(lines []string) {
	// Perform processing on the lines here
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	filename := "example.txt"

	// Using CPS
	readFileAsyncCPS(filename, func(err error, lines []string) {
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		processLines(lines)
	})

	// Using async/await (Go 1.15+)
	lines, err := readFileAsyncAwait(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	processLines(lines)
}
