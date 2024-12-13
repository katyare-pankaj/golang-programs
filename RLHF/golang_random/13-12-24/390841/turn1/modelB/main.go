package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ChannelMetrics keeps track of metrics for a specific channel
type ChannelMetrics struct {
	sends    int
	receives int
	sendWait sync.WaitGroup
	recvWait sync.WaitGroup
}

// Send increments send count and starts a timer
func (m *ChannelMetrics) Send() {
	m.sends++
	m.sendWait.Add(1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate send blocking
		m.sendWait.Done()
	}()
}

// Receive increments receive count and starts a timer
func (m *ChannelMetrics) Receive() {
	m.receives++
	m.recvWait.Add(1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate receive blocking
		m.recvWait.Done()
	}()
}

// PrintMetrics prints the channel metrics
func (m *ChannelMetrics) PrintMetrics() {
	sendWait := m.sendWait
	recvWait := m.recvWait

	sendWait.Wait()
	recvWait.Wait()

	fmt.Printf("Channel Metrics:\n")
	fmt.Printf("  Sends: %d\n", m.sends)
	fmt.Printf("  Receives: %d\n", m.receives)
}

func main() {
	var ch = make(chan int, 1) // Buffered channel with size 1
	var metrics = ChannelMetrics{}

	// Simulate goroutine activity
	const numGoroutines = 10
	const iterations = 100

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				metrics.Send()
				ch <- i // Simulate sending a value

				metrics.Receive()
				<-ch // Simulate receiving a value
			}
		}()
	}

	time.Sleep(2 * time.Second) // Allow goroutines to complete their work

	metrics.PrintMetrics()
}
