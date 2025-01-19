package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type Data struct {
	Value [1024]byte // Large struct to simulate heavy memory usage
}

func processData(count int) {
	var data []Data
	for i := 0; i < count; i++ {
		data = append(data, Data{})
	}
	// Simulate some working delay
	time.Sleep(2 * time.Second)
}

func main() {
	// Create a heap profile
	f, err := os.Create("heap.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	// Disabling garbage collection to make profiling easier
	runtime.GC()

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	processData(10000)

	fmt.Println("Processing completed")
}
