package main

import (
        "bufio"
        "fmt"
        "io"
        "log"
        "os"
        "strings"
)

// Mapper function: takes a line of text and returns a list of (word, 1) pairs
func mapper(line string) []map[string]int {
        words := strings.Fields(line)
        return make([]map[string]int, len(words))
        for i, word := range words {
                result[i] = map[string]int{word: 1}
        }
        return result
}

// Reducer function: takes a list of (word, 1) pairs and returns a single (word, count) pair
func reducer(pairs []map[string]int) map[string]int {
        counts := make(map[string]int)
        for _, pair := range pairs {
                for word, count := range pair {
                        counts[word] += count
                }
        }
        return counts
}

// Apply MapReduce to the input file and return the final word counts
func mapReduce(inputFile string) map[string]int {
        file, err := os.Open(inputFile)
        if err != nil {
                log.Fatalf("Error opening file: %v", err)
        }
        defer file.Close()

        var result map[string]int
        // Read the file line by line and apply the mapper function
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                mappedPairs := mapper(line)
                // Reduce the mapped pairs to get the final counts
                result = reducer(append(result, mappedPairs...))
        }

        if err := scanner.Err(); err != nil {
                log.Fatalf("Error scanning file: %v", err)
        }
        return result
}

func main() {
        inputFile := "input.txt" // Replace with your actual input file path
        wordCounts := mapReduce(inputFile)

        // Test the result
        expectedCounts := map[string]int{"hello": 2, "world": 1, "how": 1, "are": 1, "you": 1}
        if !cmp.Equal(wordCounts, expectedCounts) {
                log.Fatalf("Test failed: Word counts do not match. Expected: %v, Got: %v", expectedCounts, wordCounts)
        }
        fmt.Println("Word counts:")
        for word, count := range wordCounts {
                fmt.Printf("%s: %d\n", word, count)
				
			