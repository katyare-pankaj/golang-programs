package main

import (
	"fmt"
	"os"
)

func main() {
	files := make([]*os.File, 3)
	for i := range files {
		file, err := os.Create(fmt.Sprintf("file%d.txt", i))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		files[i] = file
		defer file.Close() // Schedule file closure on function return
	}

	// Perform operations with the files
	for _, file := range files {
		_, err := file.WriteString("Some data\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Wrote data to file successfully.")
	}
}
