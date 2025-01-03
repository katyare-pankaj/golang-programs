package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

const filePath = "data.txt"

func main() {
	// Create a WaitGroup to ensure all tasks complete
	var wg sync.WaitGroup

	// Initialize file or create it if it doesn't exist
	initFile()

	// Perform concurrent read and write operations
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go concurrentReadWrite(wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Read the final contents of the file
	finalContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading final file content:", err)
		return
	}

	fmt.Println("Final file content:", string(finalContent))
}

func concurrentReadWrite(wg sync.WaitGroup) {
	defer wg.Done()

	// Simulate read and write operations
	readAndWriteData()

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func readAndWriteData() {
	// Read data from the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Display the read content
	fmt.Println("Read content:", string(content))

	// Append some random data to the file
	newContent := content
	newContent = append(newContent, []byte(" Appended some data...\n")...)

	// Write new content back to the file
	err = ioutil.WriteFile(filePath, newContent, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func initFile() {
	// Check if file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create file if it doesn't exist
			err = ioutil.WriteFile(filePath, []byte("Initial file content...\n"), 0644)
			if err != nil {
				fmt.Println("Error initializing file:", err)
				return
			}
		} else {
			fmt.Println("Error checking file existence:", err)
			return
		}
	}
}
