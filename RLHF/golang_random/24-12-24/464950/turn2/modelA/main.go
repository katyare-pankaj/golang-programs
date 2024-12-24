package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
)

func main() {

	// Create a file to save the CPU profile
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile:", err)
		return
	}
	defer cpuFile.Close()

	// Start the CPU profile
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("Error starting CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	f, err := os.CreateTemp("", "example")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	content := "Hello, World!"
	if _, err := f.WriteString(content); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	readContent, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Read content:", string(readContent))
}
