package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Immutable data structure to represent a stock price
type StockPrice struct {
	Symbol string
	Price  float64
	Time   time.Time
}

// Lazy-evaluated data structure to hold a collection of stock prices
type StockPriceBook struct {
	mu        sync.Mutex
	prices    map[string]StockPrice
	lazyStats *lazyStats
}

type lazyStats struct {
	maxPrice float64
	minPrice float64
}

// Create a new StockPriceBook instance
func NewStockPriceBook() *StockPriceBook {
	return &StockPriceBook{
		prices: make(map[string]StockPrice),
	}
}

// Add a new stock price to the Book
func (spb *StockPriceBook) AddPrice(price StockPrice) {
	spb.mu.Lock()
	defer spb.mu.Unlock()

	spb.prices[price.Symbol] = price
	spb.lazyStats = nil // Invalidate the lazy stats when adding a new price
}

// Get the stock price by its symbol
func (spb *StockPriceBook) GetPrice(symbol string) (StockPrice, bool) {
	spb.mu.Lock()
	defer spb.mu.Unlock()

	price, ok := spb.prices[symbol]
	return price, ok
}

// Calculate and return the maximum price across all stocks (lazy evaluation)
func (spb *StockPriceBook) GetMaxPrice() float64 {
	spb.mu.Lock()
	defer spb.mu.Unlock()

	if spb.lazyStats == nil {
		spb.lazyStats = &lazyStats{}
		for _, price := range spb.prices {
			if spb.lazyStats.maxPrice == 0 || price.Price > spb.lazyStats.maxPrice {
				spb.lazyStats.maxPrice = price.Price
			}
			if spb.lazyStats.minPrice == 0 || price.Price < spb.lazyStats.minPrice {
				spb.lazyStats.minPrice = price.Price
			}
		}
	}
	return spb.lazyStats.maxPrice
}

// Simulate real-time data processing by adding random stock prices
func simulateRealTimeDataProcessing(spb *StockPriceBook) {
	rand.Seed(time.Now().UnixNano())
	symbols := []string{"AAPL", "MSFT", "GOOG", "AMZN", "FB"}

	for {
		symbol := symbols[rand.Intn(len(symbols))]
		price := float64(rand.Intn(1000)) + 0.99
		spb.AddPrice(StockPrice{Symbol: symbol, Price: price, Time: time.Now()})
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	spb := NewStockPriceBook()
	go simulateRealTimeDataProcessing(spb)

	for {
		time.Sleep(time.Second * 2)
		fmt.Printf("Max Price: %.2f\n", spb.GetMaxPrice())

		// Example of deep equality comparison for immutable data structures
		price1 := StockPrice{Symbol: "ABC", Price: 100.0, Time: time.Now()}
		price2 := StockPrice{Symbol: "ABC", Price: 100.0, Time: time.Now()}
		if cmp.Equal(price1, price2) {
			fmt.Println("Prices are equal.")
		}
	}
}
