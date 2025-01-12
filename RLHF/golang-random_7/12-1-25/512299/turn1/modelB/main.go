package main

import (
	"fmt"
	"math"
)

// Define the GeospatialPoint struct
type GeospatialPoint struct {
	Latitude  float64
	Longitude float64
}

// Function to calculate the distance between two geospatial points using the Haversine formula
func haversineDistance(point1, point2 GeospatialPoint) float64 {
	const earthRadius = 6371 // Radius of Earth in kilometers
	deltaLat := point2.Latitude - point1.Latitude
	deltaLon := point2.Longitude - point1.Longitude
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(point1.Latitude)*math.Cos(point2.Latitude)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

func main() {
	const threshold = 10.0 // Proximity threshold in kilometers

	// Define a slice of geospatial points
	points := []GeospatialPoint{
		{32.6514, -161.4333},
		{37.7747, -122.4182},
		{31.2345, -101.5361},
		{37.7747, -122.4172}, // A point close to the first one
		{38.8977, -77.0365},
		{32.6514, -161.4343}, // Another point close to the first one
	}

	// Create a map to store the groups of geospatial points based on their proximity
	groups := make(map[string][]GeospatialPoint)

	for _, point := range points {
		key := fmt.Sprintf("%f_%f", point.Latitude, point.Longitude)
		found := false

		// Iterate over existing groups to find the closest group
		for groupKey, groupPoints := range groups {
			for _, groupPoint := range groupPoints {
				distance := haversineDistance(point, groupPoint)
				if distance <= threshold {
					// Point is within the proximity threshold of the existing group
					groups[groupKey] = append(groups[groupKey], point)
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			// Point doesn't belong to any existing group, create a new group
			groups[key] = []GeospatialPoint{point}
		}
	}

	// Display the grouped geospatial points
	for groupKey, groupPoints := range groups {
		fmt.Printf("Group %s:\n", groupKey)
		for _, point := range groupPoints {
			fmt.Printf("(%f, %f)\n", point.Latitude, point.Longitude)
		}
		fmt.Println()
	}
}
