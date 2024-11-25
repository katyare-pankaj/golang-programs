package main

import (
	"testing"
)

type channelType int

const (
	bufferedChannel channelType = iota
	unbufferedChannel
)

func sendValue(t *testing.T, c channelType, val int) {

	var ch chan int
	switch c {
	case bufferedChannel:
		ch = make(chan int, 1)
	case unbufferedChannel:
		ch = make(chan int)
	default:
		t.Fatalf("Invalid channel type: %d", c)
	}

	go func() {
		ch <- val
		close(ch)
	}()

	received := <-ch
	if received != val {
		t.Fatalf("Expected to receive %d, but received %d", val, received)
	}
}

func TestChannelSendValue(t *testing.T) {
	testCases := []struct {
		name     string
		chanType channelType
		val      int
	}{
		{
			name:     "Send value to buffered channel",
			chanType: bufferedChannel,
			val:      100,
		},
		{
			name:     "Send value to unbuffered channel",
			chanType: unbufferedChannel,
			val:      200,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			sendValue(t, tt.chanType, tt.val)
		})
	}
}

func BenchmarkChannelSendValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var ch chan int
		val := i % 100

		// Replace the test code with the actual benchmark code
		switch i % 2 {
		case 0:
			ch = make(chan int, 1)
		default:
			ch = make(chan int)
		}

		go func() {
			ch <- val
			close(ch)
		}()

		received := <-ch
		if received != val {
			b.Fatalf("Expected to receive %d, but received %d", val, received)
		}
	}
}
