package main

import (
	"testing"
)

func BenchmarkSendToBufferedChannel(b *testing.B) {
	var ch chan int = make(chan int, 100)

	for i := 0; i < b.N; i++ {
		go func() {
			ch <- i
		}()
		<-ch
		close(ch)
	}
}

func BenchmarkSendToUnbufferedChannel(b *testing.B) {
	var ch chan int = make(chan int)

	for i := 0; i < b.N; i++ {
		go func() {
			ch <- i
		}()
		<-ch
		close(ch)
	}
}
