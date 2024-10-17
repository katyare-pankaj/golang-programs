package main

import (
	"fmt"
	"sort"
	"time"
)

// Shipment represents a shipment entry
type Shipment struct {
	Timestamp time.Time
	ProductID string
	Quantity  int
}

func aggregateDailyShipments(shipments []Shipment) map[time.Time]int {
	dailyAggregates := make(map[time.Time]int)

	// Group shipments by day
	for _, shipment := range shipments {
		day := shipment.Timestamp.Truncate(24 * time.Hour)
		dailyAggregates[day] += shipment.Quantity
	}

	return dailyAggregates
}

func processDailyAggregates(dailyAggregates map[time.Time]int) {
	for day, totalQuantity := range dailyAggregates {
		fmt.Printf("Date: %s, Total Shipped Quantity: %d\n", day.Format("2006-01-02"), totalQuantity)
	}
}

func main() {
	// Sample shipment data
	shipments := []Shipment{
		{Timestamp: time.Now().AddDate(0, 0, -2), ProductID: "P1", Quantity: 100},
		{Timestamp: time.Now().AddDate(0, 0, -2), ProductID: "P2", Quantity: 50},
		{Timestamp: time.Now().AddDate(0, 0, -1), ProductID: "P1", Quantity: 80},
		{Timestamp: time.Now().AddDate(0, 0, -1), ProductID: "P3", Quantity: 30},
		// Add more shipment data as needed
	}

	// Sort shipments by timestamp
	sort.Slice(shipments, func(i, j int) bool { return shipments[i].Timestamp.Before(shipments[j].Timestamp) })

	// Aggregate shipments by day
	dailyAggregates := aggregateDailyShipments(shipments)

	// Process the aggregated data
	fmt.Println("Daily Shipment Aggregation:")
	processDailyAggregates(dailyAggregates)
}
