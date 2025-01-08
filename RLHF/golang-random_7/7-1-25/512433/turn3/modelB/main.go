package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// MarketData represents stock price data
type MarketData struct {
    Symbol string
    Price  float64
    Time   time.Time
}

// MarketMonitor monitors real-time market data and generates insights
type MarketMonitor struct {
    dataCh     chan MarketData
    signalCh   chan string
    stopCh     chan struct{}
    wg         sync.WaitGroup
    ticker     *time.Ticker
    priceThreshold float64
}

// NewMarketMonitor creates a new MarketMonitor instance
func NewMarketMonitor(priceThreshold float64) *MarketMonitor {
    return &MarketMonitor{
        dataCh:     make(chan MarketData, 100),
        signalCh:   make(chan string, 100),
        stopCh:     make(chan struct{}),
        ticker:     time.NewTicker(5 * time.Second),
        priceThreshold: priceThreshold,
    }
}

// Start starts the monitoring system
func (m *MarketMonitor) Start() {
    m.wg.Add(1)
    go m.monitor()
}

// Stop stops the monitoring system
func (m *MarketMonitor) Stop() {
    close(m.stopCh)
    m.wg.Wait()
}

// IngestData ingests market data into the monitoring system
func (m *MarketMonitor) IngestData(data MarketData) {
    m.dataCh <- data
}

// GetSignalChannel returns the channel to receive signals
func (m *MarketMonitor) GetSignalChannel() <-chan string {
    return m.signalCh
}

func (m *MarketMonitor) monitor() {
    defer m.wg.Done()

    for {
        select {
        case data := <-m.dataCh:
            m.handleData(data)
        case <-m.ticker.C:
            m.generateInsights()
        case <-m.stopCh:
            m.ticker.Stop()
            close(m.signalCh)
            return
        }
    }
}

func (m *MarketMonitor) handleData(data MarketData) {
    // Perform any necessary data processing and store it if required
    // For simplicity, we'll just print the received data
    fmt.Printf("Received data: %+v\n", data)
}

func (m *MarketMonitor) generateInsights() {
    // Generate insights based on the processed data
    // In this example, we'll just generate a random signal
    if rand.Float64() > 0.5 {
        m.signalCh <- "BUY"
    } else {
        m.signalCh <- "SELL"
    }
}

// DataFetcher simulates fetching market data
type DataFetcher struct {
    symbols     []string
    dataCh      chan MarketData
    stopCh      chan struct{}
    wg          sync.WaitGroup
    fetchInterval time.Duration
}

// NewDataFetcher creates a new DataFetcher instance
func NewDataFetcher(symbols []string, fetchInterval time.Duration) *DataFetcher {
    return &DataFetcher{
        symbols:     symbols,
        dataCh:      make(chan MarketData, 100),
        stopCh:      make(chan struct{}),
        fetchInterval: fetchInterval,
    }
}

// Start starts the data fetching system
func (f *DataFetcher) Start() {
    f.wg.Add(1)
    go f.fetchData()
}

// Stop stops the data fetching system
func (f *DataFetcher) Stop() {
    close(f.stopCh)
    f.wg.Wait()
}

// GetDataChannel returns the channel to receive market data
func (f *DataFetcher) GetDataChannel() <-chan MarketData {
    return f.dataCh
}

func (f *DataFetcher) fetchData() {
    defer f.wg.Done()

    ticker := time.NewTicker(f.fetchInterval)