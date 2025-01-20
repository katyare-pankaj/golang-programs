package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

type row struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Simulate a memory leak by keeping a reference to an unnecessary slice
	var leakedRows []*row
	for i := 0; i < 10000; i++ {
		leakedRows = append(leakedRows, &row{})
	}
	defer func() {
		// Clear the leaked rows to free the memory
		leakedRows = nil
	}()
	// Open a CSV file and read its content
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		r := row{
			Name:  record[0],
			Age:   0,
			Email: record[2],
		}
		// Use the range loop to print the row values
		for _, value := range r {
			fmt.Printf("%s ", value)
		}
		fmt.Println()
	}
	// Check for memory leaks using the debug package
	stats := &debug.MemStats{}
	debug.ReadMemStats(stats)
	allocated := stats.Alloc
	if allocated > 10000000 { // 10MB in bytes
		fmt.Println("Possible memory leak detected!")
		fmt.Println("Memory allocated:", allocated/1024/1024, "MB")
	}
	// Wait for a bit to see if the memory leak is cleared
	time.Sleep(time.Second * 5)
	debug.ReadMemStats(stats)
	allocated = stats.Alloc
	fmt.Println("Memory allocated after sleep:", allocated/1024/1024, "MB")
}
