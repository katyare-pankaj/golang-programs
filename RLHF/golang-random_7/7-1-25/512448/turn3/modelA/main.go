package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// StockPrice struct for storing stock prices
type StockPrice struct {
	Ticker string
	Price  float64
	Time   time.Time
}

// FetchData simulates fetching stock prices from a ticker
func fetchData(priceCh chan<- StockPrice, ticker string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		price := rand.Float64()*100.0 + float64(ticker[0]-'A')*100.0 // Sample price based on ticker
		data := StockPrice{Ticker: ticker, Price: price, Time: time.Now()}
		priceCh <- data
		time.Sleep(time.Second) // Simulate fetch interval
	}
	close(priceCh)
}

// CalculateMovingAverage calculates a simple moving average
func calculateMovingAverage(priceCh <-chan StockPrice, maCh chan<- float64, ticker string, window int) {
	var prices []float64
	for price := range priceCh {
		if price.Ticker != ticker {
			continue
		}
		prices = append(prices, price.Price)
		if len(prices) > window {
			prices = prices[1:]
		}
		if len(prices) >= window {
			maCh <- sum(prices) / float64(window)
		}
	}
	close(maCh)
}

// Sum calculates the sum of a slice of floats
func sum(nums []float64) float64 {
	var total float64
	for _, num := range nums {
		total += num
	}
	return total
}

// DetectPriceChange detects significant price changes
func detectPriceChange(maCh <-chan float64, changeCh chan<- string, threshold float64) {
	var previous float64
	for ma := range maCh {
		if ma < previous-threshold || ma > previous+threshold {
			changeCh <- fmt.Sprintf("Significant change: %.2f", ma)
		}
		previous = ma
	}
	close(changeCh)
}

func main() {
	var wg sync.WaitGroup

	tickers := []string{"AAPL", "GOOGL", "MSFT"}
	priceChannels := make([]chan StockPrice, len(tickers))
	maChannels := make([]chan float64, len(tickers))
	changeChannels := make([]chan string, len(tickers))

	// Set up channels and start fetching data
	for i, ticker := range tickers {
		priceChannels[i] = make(chan StockPrice)
		maChannels[i] = make(chan float64)
		changeChannels[i] = make(chan string)

		wg.Add(1)
		go fetchData(priceChannels[i], ticker, &wg)

		wg.Add(1)
		go calculateMovingAverage(priceChannels[i], maChannels[i], ticker, 10)

		wg.Add(1)
		go detectPriceChange(maChannels[i], changeChannels[i], 5.0)
	}

	// Monitor change notifications
	for _, changeCh := range changeChannels {
		go func(ch <-chan string) {
			for change := range ch {
				fmt.Println(change)
			}
		}(changeCh)
	}

	wg.Wait()
	fmt.Println("System terminated.")
}
