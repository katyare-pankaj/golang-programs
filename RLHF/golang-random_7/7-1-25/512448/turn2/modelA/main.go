package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// StockData holds stock price data
type StockData struct {
	Symbol string
	Price  float64
	Time   time.Time
}

// fetchStockData simulates fetching stock data streams
func fetchStockData(stockDataCh chan<- StockData, symbol string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		price := rand.Float64()*100.0 + 10.0 // between 10 and 110
		data := StockData{Symbol: symbol, Price: price, Time: time.Now()}
		stockDataCh <- data
		time.Sleep(time.Second) // Simulate delay in data arrival
	}
	close(stockDataCh)
}

// calculateSignal calculates trading signal based on price
func calculateSignal(data ChannelOfStockData) <-chan string {
	signalCh := make(chan string)

	go func() {
		for data := range data {
			// Implement real logic here (e.g., moving averages, indicators)
			if data.Price > 105 {
				signalCh <- "SELL"
			} else {
				signalCh <- "BUY"
			}
		}
		close(signalCh)
	}()

	return signalCh
}

type ChannelOfStockData <-chan StockData

func main() {
	var wg sync.WaitGroup

	// Symbol channels for multiple stocks
	symbols := []string{"AAPL", "GOOGL", "MSFT"}
	stockDataChannels := make([]ChannelOfStockData, len(symbols))

	// Initialize channels and start data fetching goroutines
	for i, symbol := range symbols {
		stockDataCh := make(chan StockData)
		stockDataChannels[i] = stockDataCh

		wg.Add(1)
		go fetchStockData(stockDataCh, symbol, &wg)
	}

	// Signal channels for calculated signals
	signalChannels := make([]chan string, len(symbols))

	for i, dataCh := range stockDataChannels {
		signalChannels[i] = calculateSignal(dataCh)
	}

	// Monitor signals
	for i, signalCh := range signalChannels {
		fmt.Printf("Monitoring stock %s:\n", symbols[i])
		for signal := range signalCh {
			fmt.Println(signal)
		}
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}
