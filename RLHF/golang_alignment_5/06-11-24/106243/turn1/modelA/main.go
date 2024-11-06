package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Start the CPU profile
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// Simulate some work with potential bottlenecks
	doSomeWork()

	// Memory profiling
	memf, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer memf.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	fmt.Println("Performance analysis completed.")
}

func doSomeWork() {
	const numIterations = 1000000
	sum := 0

	// Bottlenecked loop
	for i := 0; i < numIterations; i++ {
		sum += expensiveFunction(i)
	}
	fmt.Println("Sum:", sum)
}

func expensiveFunction(x int) int {
	// Simulate an expensive operation
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	return x * x
}
