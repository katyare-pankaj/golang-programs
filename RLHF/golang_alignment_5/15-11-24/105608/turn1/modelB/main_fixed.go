package main

import (
	"fmt"
	"sync"
	"time"
)

// Well represents the abstraction of an oil well
type Well struct {
	name     string
	capacity int
}

// Pump represents a pump for extracting oil from a well.
type Pump struct {
	well    *Well
	maxRate int
	work    chan PumpWork // channel for work requests
	wg      *sync.WaitGroup
}

// PumpWork represents a unit of work for a pump
type PumpWork struct {
	Amount int
}

// start pumps the oil from the well.
func (p *Pump) start() {
	defer p.wg.Done()
	for work := range p.work {
		if work.Amount <= 0 {
			continue
		}
		// Ensure we don't pump more than the max rate
		if work.Amount > p.maxRate {
			work.Amount = p.maxRate
		}
		// Call extractOil to simulate the pumping process
		p.well.extractOil(work.Amount)
	}
}

// extractOil simulates extracting oil from the well
func (w *Well) extractOil(amount int) {
	fmt.Printf("Well[%s]: Extracting %d barrels\n", w.name, amount)
	// Simulate the extraction time being proportional to the amount
	time.Sleep(time.Duration(amount) * 1 * time.Millisecond)
}

func main() {
	well1 := Well{name: "Well_A", capacity: 100000}
	pump1 := Pump{
		well:    &well1,
		maxRate: 2000,
		work:    make(chan PumpWork, 100),
		wg:      &sync.WaitGroup{},
	}

	// Creating a workload
	workload := []PumpWork{
		{1000}, {500}, {3000}, {800}, {1500},
	}

	// Set up the WaitGroup for goroutine synchronization
	pump1.wg.Add(1)

	// Start the pump in a goroutine
	go pump1.start()

	// Serve workload by sending work to the pump's work channel
	for _, work := range workload {
		pump1.work <- work
	}

	// Close the work channel to signal the pump to finish
	close(pump1.work)

	// Wait for all work to be completed
	pump1.wg.Wait()

	fmt.Println("All work completed.")
}
