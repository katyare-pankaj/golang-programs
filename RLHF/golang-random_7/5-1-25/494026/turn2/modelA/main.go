package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var (
	wg      sync.WaitGroup
	counter int64
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	wg.Add(1)
	defer wg.Done()

	time.Sleep(time.Duration(r.FormValue("delay")) * time.Millisecond)

	counter++
	fmt.Fprintln(w, "Response:", counter)
}

func main() {
	// Start memory profiling
	memProfileFile, _ := os.Create("mem.prof")
	defer memProfileFile.Close()
	runtime.GC() // Start with a clean GC slate
	pprof.WriteHeapProfile(memProfileFile)

	// Set number of workers for handling HTTP requests
	numWorkers := runtime.NumCPU() * 2
	runtime.GOMAXPROCS(numWorkers)
	fmt.Printf("Running with %d workers.\n", numWorkers)

	// Serve HTTP requests concurrently
	http.HandleFunc("/", handleRequest)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, "Server error:", err)
		os.Exit(1)
	}
}
