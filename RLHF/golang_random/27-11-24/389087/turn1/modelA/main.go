package main

import (
	"fmt"
	"sync"
	"time"
)

// Subscriber interface to handle received data.
type Subscriber interface {
	OnNext(int)
	OnError(error)
	OnComplete()
}

// Observable interface to subscribe and receive data.
type Observable interface {
	Subscribe(Subscriber)
}

// simpleObservable implements the Observable interface.
type simpleObservable struct {
	ch      chan int
	wg      sync.WaitGroup
	errs    chan error
	done    chan bool
	sub     Subscriber
	closing bool
}

func NewSimpleObservable() *simpleObservable {
	return &simpleObservable{
		ch:   make(chan int),
		errs: make(chan error),
		done: make(chan bool),
	}
}

func (o *simpleObservable) Subscribe(sub Subscriber) {
	if o.closing {
		sub.OnError(fmt.Errorf("observable is closed"))
		return
	}

	o.sub = sub

	go func() {
		defer o.wg.Done()
		for {
			select {
			case data, ok := <-o.ch:
				if !ok {
					o.sub.OnComplete()
					return
				}
				o.sub.OnNext(data)
			case err := <-o.errs:
				o.sub.OnError(err)
				return
			case <-o.done:
				o.sub.OnComplete()
				return
			}
		}
	}()
	o.wg.Add(1)
}

func (o *simpleObservable) OnNext(data int) {
	select {
	case o.ch <- data:
	default:
	}
}

func (o *simpleObservable) OnError(err error) {
	select {
	case o.errs <- err:
	default:
	}
}

func (o *simpleObservable) OnComplete() {
	select {
	case o.done <- true:
	default:
	}
}

func (o *simpleObservable) Close() {
	o.closing = true
	o.OnComplete()
	o.wg.Wait()
}

func main() {
	obs := NewSimpleObservable()

	// Define a subscriber
	type exampleSubscriber struct {
		counter int
	}

	sub := &exampleSubscriber{}

	sub.OnNext = func(data int) {
		sub.counter++
		fmt.Println("Received data:", data, "Counter:", sub.counter)
	}

	sub.OnError = func(err error) {
		fmt.Println("Error:", err)
	}

	sub.OnComplete = func() {
		fmt.Println("Observable completed.")
	}

	obs.Subscribe(sub)

	// Emit some data
	go func() {
		for i := 1; i <= 5; i++ {
			obs.OnNext(i)
		}
		obs.OnComplete()
	}()

	// Simulate a delay for demonstration
	time.Sleep(time.Second)

	// Close the observable
	obs.Close()
}
