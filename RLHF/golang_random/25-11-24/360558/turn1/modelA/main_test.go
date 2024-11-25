package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelUseCases(t *testing.T) {
	// Define a test function that uses channel parameter(s)
	testChannelFunction := func(t *testing.T, bufferSize int, concurrentSend bool) {
		// Create a channel with the specified buffer size
		ch := make(chan int, bufferSize)

		// Optionally start goroutines for concurrent operations
		if concurrentSend {
			go func() {
				for i := 0; i < 10; i++ {
					ch <- i
					time.Sleep(time.Millisecond * 1)
				}
			}()
		}

		// Simulate send operations
		for i := 0; i < 10; i++ {
			if concurrentSend {
				time.Sleep(time.Millisecond * 1)
			}
			ch <- i
		}

		// Simulate receive operations
		var receivedValues []int
		for i := 0; i < 10; i++ {
			val := <-ch
			receivedValues = append(receivedValues, val)
			if concurrentSend {
				time.Sleep(time.Millisecond * 1)
			}
		}

		// Verify received values are as expected
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, receivedValues)
	}

	// Define test cases with different parameters
	testCases := []struct {
		bufferSize     int
		concurrentSend bool
	}{
		{bufferSize: 0, concurrentSend: false}, // Unbuffered channel, non-concurrent
		{bufferSize: 5, concurrentSend: false}, // Buffered channel, non-concurrent
		{bufferSize: 0, concurrentSend: true},  // Unbuffered channel, concurrent
		{bufferSize: 5, concurrentSend: true},  // Buffered channel, concurrent
	}

	// Run test function with each parameter set
	for _, testCase := range testCases {
		t.Run("bufferSize=%d,concurrentSend=%t", func(t *testing.T) {
			testChannelFunction(t, testCase.bufferSize, testCase.concurrentSend)
		})
	}
}
