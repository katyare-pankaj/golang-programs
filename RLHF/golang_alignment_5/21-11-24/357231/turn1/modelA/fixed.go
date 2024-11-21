package main

import (
	"fmt"
	"log"
	"net/http"

	"time"
)

func busyLoop() {
	for {
	}
}

func main() {
	go busyLoop()               // Simulating some work
	time.Sleep(time.Second * 5) // Sleep for 5 seconds to allow profiling

	fmt.Println("Starting HTTP server for profiling...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profiling available at /debug/pprof/")
	})

	log.Fatal(http.ListenAndServe(":6060", nil))
}
