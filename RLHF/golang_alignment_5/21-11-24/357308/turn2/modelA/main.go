package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

var profile *os.File

func main() {
	// Command-line flag to specify the profile file
	profileFile := flag.String("profile", "mem.prof", "write the profile to this file")
	flag.Parse()

	// Start memory profiling
	if profileFile != nil && *profileFile != "" {
		var err error
		profile, err = os.Create(*profileFile)
		if err != nil {
			fmt.Printf("error opening profile file: %v\n", err)
			os.Exit(1)
		}
		defer profile.Close()
		pprof.StartHeapProfile(pprof.ProfileAllocs|pprof.ProfileInuseSpace, 1024)
		defer pprof.StopHeapProfile()
	}

	// Simulate memory allocation
	for i := 0; i < 1000000; i++ {
		make([]byte, i)
		time.Sleep(time.Microsecond)
	}

	// Dump the profile
	if profile != nil {
		pprof.WriteHeapProfile(profile)
	}
}
