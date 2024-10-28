package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

const (
	numIterations = 1000000
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func simulateErrorHandling() error {
	if rand.Intn(100) < 5 {
		return fmt.Errorf("simulated error")
	}
	return nil
}

func main() {
	// Profiling
	f, err := os.Create("profile.prof")
	if err != nil {
		log.Fatal("could not create profile: ", err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Benchmarking
	benchmarkErrorHandling()
}

func benchmarkErrorHandling() {
	for i := 0; i < numIterations; i++ {
		err := simulateErrorHandling()
		handleError(err)
	}
}

func TestErrorHandlingPerformance(t *testing.T) {
	start := time.Now()
	benchmarkErrorHandling()
	duration := time.Since(start)

	t.Logf("Execution time: %s", duration)
	// Add assertions to check performance thresholds if required
}
