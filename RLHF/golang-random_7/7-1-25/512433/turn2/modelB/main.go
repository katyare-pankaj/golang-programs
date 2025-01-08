package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MarketData represents a single price update from a financial data source.
type MarketData struct {
	Symbol string
	Price  float64
	Time   time.Time
	Volume int
}

// dataStreamGenerator generates a continuous stream of market data.
func dataStreamGenerator(symbol string, dataCh chan<- MarketData, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Generate random market data with a slight delay
		price := rand.Float64() * 100.0
		volume := rand.Intn(1000)
		dataCh <- MarketData{Symbol: symbol, Price: price, Time: time.Now(), Volume: volume}
		time.Sleep(time.Millisecond * 100)
	}
}

// simpleMovingAverage calculates the simple moving average for a given window size.
func simpleMovingAverage(symbol string, windowSize int, dataCh <-chan MarketData, avgCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	window := make([]float64, windowSize)
	var sum float64

	for data := range dataCh {
		if len(window) == windowSize {
			// Calculate the average and remove the oldest value from the window
			sum -= window[0]
			window = window[1:]
		}
		window = append(window, data.Price)
		sum += data.Price
		avgCh <- sum / float64(len(window))
	}
}

// trader makes decisions based on the moving average and executes trades.
func trader(symbol string, avgCh <-chan float64, tradeCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	const (
		buyingThreshold  = 80.0
		sellingThreshold = 20.0
	)
	for avg := range avgCh {
		if avg >= buyingThreshold {
			tradeCh <- fmt.Sprintf("%s: Buy at %.2f", symbol, avg)
		} else if avg <= sellingThreshold {
			tradeCh <- fmt.Sprintf("%s: Sell at %.2f", symbol, avg)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	dataCh := make(chan MarketData)
	avgCh := make(chan float64)
	tradeCh := make(chan string)

	// Generate data stream for multiple symbols
	symbols := []string{"GOOG", "AAPL", "MSFT"}
	for _, symbol := range symbols {
		wg.Add(1)
		go dataStreamGenerator(symbol, dataCh, &wg)
	}

	// Calculate moving average for each symbol concurrently
	for _, symbol := range symbols {
		wg.Add(1)
		go simpleMovingAverage(symbol, 20, dataCh, avgCh, &wg)
	}

	// Execute trades based on the moving average for each symbol concurrently
	for _, symbol := range symbols {
		wg.Add(1)
		go trader(symbol, avgCh, tradeCh, &wg)
	}

	// Receive and display trades
	go func() {
		for trade := range tradeCh {
			fmt.Println(trade)
		}
	}()

	wg.Wait()
	fmt.Println("Algorithm execution completed.")
}
