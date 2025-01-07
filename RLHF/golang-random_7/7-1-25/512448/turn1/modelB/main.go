package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MarketData struct {
	Price float64
	Time  time.Time
}

func dataGenerator(dataCh chan<- MarketData, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		price := rand.Float64() * 100.0
		dataCh <- MarketData{Price: price, Time: time.Now()}
		time.Sleep(time.Millisecond * 5)
	}
	close(dataCh)
}

func adaptiveAlgorithm(dataCh <-chan MarketData, wg *sync.WaitGroup) {
	defer wg.Done()
	var lastPrice float64
	for data := range dataCh {
		// Simple adaptive algorithm: adjust threshold based on price change
		threshold := 0.5
		if data.Price > lastPrice+threshold {
			fmt.Println("Price increase detected:", data.Price)
		} else if data.Price < lastPrice-threshold {
			fmt.Println("Price decrease detected:", data.Price)
		}
		lastPrice = data.Price
	}
}

func main() {
	var wg sync.WaitGroup
	dataCh := make(chan MarketData)

	wg.Add(2)
	go dataGenerator(dataCh, &wg)
	go adaptiveAlgorithm(dataCh, &wg)

	wg.Wait()
	fmt.Println("Algorithm execution completed.")
}
