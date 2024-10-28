package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
)

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func divide(a, b int) error {
	if b == 0 {
		return fmt.Errorf("division by zero")
	}
	return nil
}

func doWork() error {
	a := rand.Intn(100)
	b := rand.Intn(100)
	err := divide(a, b)
	handleError(err)
	return err
}

func BenchmarkDoWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWork()
	}
}

func main() {
	// Profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// Benchmarking
	testing.Benchmark(BenchmarkDoWork)
}
