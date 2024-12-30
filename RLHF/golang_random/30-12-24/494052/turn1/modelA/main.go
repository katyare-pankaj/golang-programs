package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func processHeritageData(inputFile, outputFile string) error {
	// Open input file
	inputReader, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputReader.Close()

	// Create CSV reader
	csvReader := csv.NewReader(bufio.NewReader(inputReader))

	// Open output file
	outputWriter, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputWriter.Close()

	// Create CSV writer
	csvWriter := csv.NewWriter(outputWriter)

	// Read header
	header, err := csvReader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	// Write header to output
	if err := csvWriter.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Process rows
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			return fmt.Errorf("failed to read row: %w", err)
		}

		year, err := strconv.Atoi(row[2])
		if err != nil {
			log.Printf("Warning: Invalid year '%s' for row %s, skipping", row[2], row)
			continue // Skip row with invalid year
		}

		if year >= 1900 {
			if err := csvWriter.Write(row); err != nil {
				return fmt.Errorf("failed to write row: %w", err)
			}
		}
	}

	// Flush the writer
	if err := csvWriter.Flush(); err != nil {
		return fmt.Errorf("failed to flush output writer: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <input_file> <output_file>", os.Args[0])
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := processHeritageData(inputFile, outputFile)
	if err != nil {
		log.Fatalf("Processing failed: %v", err)
	}

	log.Printf("Processing completed successfully.")
}
