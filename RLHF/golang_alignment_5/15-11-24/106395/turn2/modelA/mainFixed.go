package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

const (
	fileName  = "data.csv"
	chunkSize = 1000 // Process 1000 rows at a time
)

type dataChunk struct {
	rows [][]string
}

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func splitData(data [][]string) []dataChunk {
	var chunks []dataChunk
	for i := 0; i < len(data); i += chunkSize {
		end := min(i+chunkSize, len(data))
		chunks = append(chunks, dataChunk{data[i:end]})
	}
	return chunks
}

func processChunk(chunk dataChunk, wg *sync.WaitGroup) {
	defer wg.Done()
	// Process the chunk here
	for _, row := range chunk.rows {
		fmt.Println("Processing row:", row)
		// Processing logic for each row
	}
}

func processDataConcurrently(data [][]string) {
	wg := &sync.WaitGroup{}
	chunks := splitData(data)

	for _, chunk := range chunks {
		wg.Add(1)
		go processChunk(chunk, wg)
	}

	wg.Wait()
	fmt.Println("Data processing completed.")
}

func main() {
	data, err := readCSV(fileName)
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	processDataConcurrently(data)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
