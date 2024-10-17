package main

import (
	"fmt"
)

type Shipment struct {
	TrackingNumber string
	Destination    string
	Weight         int
}

func aggregateShipments(shipments []Shipment) map[string][]Shipment {
	aggregatedShipments := make(map[string][]Shipment)
	for _, shipment := range shipments {
		destination := shipment.Destination
		if _, exists := aggregatedShipments[destination]; !exists {
			aggregatedShipments[destination] = []Shipment{shipment}
		} else {
			aggregatedShipments[destination] = append(aggregatedShipments[destination], shipment)
		}
	}
	return aggregatedShipments
}

func main() {
	shipments := []Shipment{
		{TrackingNumber: "12345", Destination: "New York", Weight: 100},
		{TrackingNumber: "67890", Destination: "New York", Weight: 50},
		{TrackingNumber: "01112", Destination: "Los Angeles", Weight: 200},
		{TrackingNumber: "34567", Destination: "San Francisco", Weight: 80},
		{TrackingNumber: "78901", Destination: "San Francisco", Weight: 120},
	}
	aggregatedShipments := aggregateShipments(shipments)
	for destination, shipments := range aggregatedShipments {
		fmt.Printf("Destination: %s\n", destination)
		for _, shipment := range shipments {
			fmt.Printf("Tracking Number: %s, Weight: %d\n", shipment.TrackingNumber, shipment.Weight)
		}
		fmt.Println()
	}
}
