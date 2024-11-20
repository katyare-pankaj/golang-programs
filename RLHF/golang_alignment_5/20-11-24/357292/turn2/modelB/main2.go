package main

import (
	"fmt"
	"time"

	"github.com/cockroachdb/elapsedtime"
)

func sleepForOneSecond() {
	<-time.After(time.Second)
}

func main() {
	et := elapsedtime.New()

	sleepForOneSecond()

	elapsed := et.Elapsed()
	fmt.Println("Elapsed time:", elapsed)
}
