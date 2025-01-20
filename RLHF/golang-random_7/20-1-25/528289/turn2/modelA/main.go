package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func readCsvFile(filename string) {
	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	reader := csv.NewReader(file)

	// To read all values and ignore size: uncomment next two lines
	// records, err := reader.ReadAll()
	// checkError(err)
	// fmt.Println(records)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError(err)

		// Print all CSV fields for currently read record.
		fmt.Println(record)
	}
}

func printMemoryUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(memStats.Sys))
	fmt.Printf("\tNumGC = %v\n", memStats.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the CSV file path as an argument.")
		os.Exit(1)
	}

	// Print initial memory usage
	fmt.Println("Initial Memory Usage:")
	printMemoryUsage()
	fmt.Println("======================")

	filename := os.Args[1]
	fmt.Printf("Reading CSV file: %s\n", filename)
	readCsvFile(filename)
	fmt.Println("======================")

	// Print memory usage after reading CSV
	fmt.Println("Memory Usage After Reading CSV:")
	printMemoryUsage()
}
