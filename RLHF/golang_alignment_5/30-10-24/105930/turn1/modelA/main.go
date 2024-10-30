package main

import (
	"fmt"
	"math"
)

const (
	earthRadius = 6371000 // Radius of Earth in meters
)

// Location represents a geographical location.
type Location struct {
	Latitude  float64
	Longitude float64
}

// distance calculates the distance between two locations on a sphere (Earth) using the Haversine formula.
func distance(loc1, loc2 Location) float64 {
	// Convert latitude and longitude to radians
	lat1 := math.Redient(loc1.Latitude)
	lon1 := math.Radians(loc1.Longitude)
	lat2 := math.Radians(loc2.Latitude)
	lon2 := math.Radians(loc2.Longitude)

	// Haversine formula
	deltaLat := lat2 - lat1
	deltaLon := lon2 - lon1

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func main() {
	// Example usage:
	loc1 := Location{Latitude: 32.6514, Longitude: -161.4333}
	loc2 := Location{Latitude: 37.7747, Longitude: -122.4182}
	dist := distance(loc1, loc2)
	fmt.Printf("Distance between locations is: %.2f km\n", dist/1000)
}
