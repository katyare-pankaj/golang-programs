package main

import (
	"fmt"
	"runtime"
)

// A simple struct to demonstrate a potential memory leak
type Data struct {
	content string
}

func printContent(d *Data) {
	fmt.Println(d.content)
}

func main() {
	var mem runtime.MemStats

	// Example 1: Memory Leak
	fmt.Println("Example 1: Memory Leak")
	dataSet := []*Data{
		{content: "First"},
		{content: "Second"},
		{content: "Third"},
	}

	fmt.Println("- While ranging over the slice, closure captures loop variable causing potential memory leak:")
	for _, data := range dataSet {
		printContent(data) // Correct usage: Pass the index
		// After each iteration, the 'data' variable changes, causing older references to stay in closure.
	}

	runtime.ReadMemStats(&mem)
	fmt.Printf("Allocated memory: %v bytes\n\n", mem.Alloc)

	// Clean any stored or held pointers to enforce GC
	runtime.GC()

	// Example 2: Avoiding Memory Leak using correct scope
	fmt.Println("Example 2: Avoiding Memory Leak")

	infoSet := []*Data{
		{content: "Fourth"},
		{content: "Fifth"},
		{content: "Sixth"},
	}

	fmt.Println("- Correct usage, by re-declaring variable inside loop, avoiding closure variable capture:")
	{
		for _, d := range infoSet {
			data := d // Capture in new scope
			printContent(data)
		}
	}

	runtime.ReadMemStats(&mem)
	fmt.Printf("Allocated memory after fixing: %v bytes\n\n", mem.Alloc)

	// Example 3: Range loop over map demonstrating capture avoidance
	fmt.Println("Example 3: Ranging over Map")

	dataMap := map[string]*Data{
		"key1": {content: "Seventh"},
		"key2": {content: "Eighth"},
		"key3": {content: "Ninth"},
	}

	for k, v := range dataMap {
		key, value := k, v  // Redeclare and avoid any incorrect captures
		printContent(value) // No memory leak, the variable is re-declared
		fmt.Printf("(For key: %s)\n", key)
	}

	runtime.ReadMemStats(&mem)
	fmt.Printf("Allocated memory after ranging map: %v bytes\n", mem.Alloc)
}
