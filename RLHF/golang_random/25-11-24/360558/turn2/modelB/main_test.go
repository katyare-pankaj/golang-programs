package main

import (
	"sync"
	"testing"
)

func BenchmarkSendToBufferedChannel(b *testing.B) {
	ch := make(chan int, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			ch <- i
			close(ch)
		}()
		<-ch
	}
}

func BenchmarkSendToUnbufferedChannel(b *testing.B) {
	ch := make(chan int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			ch <- i
			close(ch)
		}()
		<-ch
	}
}

func BenchmarkSendToBufferedChannelSynchronized(b *testing.B) {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			ch <- i
			close(ch)
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkSendToUnbufferedChannelSynchronized(b *testing.B) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			ch <- i
			close(ch)
			wg.Done()
		}()
	}
	wg.Wait()
}
