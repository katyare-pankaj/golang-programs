// main.go

package main

import (
	"fmt"
	"runtime/pprof"
	"shapes"
)

func main() {
	// Create instances of shapes for profiling
	rect := shapes.Rectangle{Length: 10, Width: 20}
	circle := shapes.Circle{Radius: 5}

	// Enable profiling
	f, err := pprof.Create("main.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Simulate some usage of the shapes
	for i := 0; i < 1000000; i++ {
		rect.Area()
		circle.Area()
	}

	fmt.Println("Profiling completed.")
}
