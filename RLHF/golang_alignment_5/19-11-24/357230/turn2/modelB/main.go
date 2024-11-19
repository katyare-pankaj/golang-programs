package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

func countWords(filename string, wordCount chan int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	words := strings.Fields(string(content))
	wordCount <- len(words)
}
func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}
	wordCountChan := make(chan int)
	var wg sync.WaitGroup
	// Start a goroutine for each file to count words
	for _, filename := range filenames {
		wg.Add(1)
		go func(f string) {
			countWords(f, wordCountChan)
			wg.Done()
		}(filename)
	}

	// Total word count will be stored here
	totalWordCount := 0

	// Receive word counts from the channel and add them to total
	go func() {
		for count := range wordCountChan {
			totalWordCount += count
		}
	}()

	wg.Wait() // Wait for all goroutines to finish

	// Close the wordCount channel to signal the end of data
	close(wordCountChan)

	fmt.Println("Total words:", totalWordCount)
}
