package main

import (
	"fmt"
	"io/ioutil"
)

func processFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the data here
	fmt.Println("File content:", string(data))
}

func main() {
	filePath := "example.txt"
	processFile(filePath)
}
