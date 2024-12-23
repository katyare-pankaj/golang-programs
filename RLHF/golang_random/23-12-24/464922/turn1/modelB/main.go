package main

import (
	"fmt"
	"sync"
	"time"
)

// Two stateful microservices
type MicroService1 struct {
	data int
}

type MicroService2 struct {
	data string
}

func (m *MicroService1) Inc() {
	m.data++
}

func (m *MicroService1) Get() int {
	return m.data
}

func (m *MicroService2) Set(s string) {
	m.data = s
}

func (m *MicroService2) Get() string {
	return m.data
}

func main() {
	var wg sync.WaitGroup
	// Create instances of the stateful microservices
	ms1 := &MicroService1{}
	ms2 := &MicroService2{}

	// Let's assume these microservices are running concurrently and communicating via channels.
	// For this example, we'll use simple function calls to demonstrate synchronization.

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Microservice 1 is doing some work and updating its state.
		for i := 0; i < 5; i++ {
			ms1.Inc()
			time.Sleep(time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Microservice 2 is also doing some work and updating its state.
		for i := 0; i < 3; i++ {
			ms2.Set(fmt.Sprintf("Update %d", i))
			time.Sleep(time.Second)
		}
	}()

	// Synchronization: Wait for both microservices to finish their work before accessing their states.
	wg.Wait()

	// Now, both microservices have finished their work. We can access their states.
	fmt.Println("MicroService1 data:", ms1.Get())
	fmt.Println("MicroService2 data:", ms2.Get())
}
