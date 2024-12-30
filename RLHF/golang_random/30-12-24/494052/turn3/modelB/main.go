package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countLines(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	content := string(data)
	return len(bufio.NewScanner(strings.NewReader(content)).Split(bufio.ScanLines)), nil
}

func main() {
	filename := "test.txt"
	lineCount, err := countLines(filename)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Number of lines in %s: %d\n", filename, lineCount)
}
