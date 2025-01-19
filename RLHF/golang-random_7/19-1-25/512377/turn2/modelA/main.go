package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func processLargeData() {
	data := make(map[string]int)

	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("key-%d", i%1000)
		data[key]++
	}

	total := 0
	for _, v := range data {
		total += v
	}

	fmt.Println("Total:", total)

}

func main() {
	// Start the pprof server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Optionally create a CPU profile file
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Run the data processing function
	processLargeData()
	time.Sleep(60 * time.Second)

}
