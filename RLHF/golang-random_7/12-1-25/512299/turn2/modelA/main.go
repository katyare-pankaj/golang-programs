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

// Function to calculate the centroid of a set of geospatial points
func calculateCentroid(points []GeospatialPoint) GeospatialPoint {
	n := float64(len(points))
	if n == 0 {
		return GeospatialPoint{} // Return empty point if no points
	}

	sumLat := 0.0
	sumLon := 0.0
	for _, point := range points {
		sumLat += point.Latitude
		sumLon += point.Longitude
	}
	return GeospatialPoint{Latitude: sumLat / n, Longitude: sumLon / n}
}

// Function to cluster GPS coordinates based on distance
func clusterCoordinates(points []GeospatialPoint, threshold float64) map[GeospatialPoint][]GeospatialPoint {
	clusters := make(map[GeospatialPoint][]GeospatialPoint)

	// Assign initial centroids to the first few points
	for i := range points[:3] {
		clusters[points[i]] = []GeospatialPoint{points[i]}
	}

	// Iterate through all points and assign them to the nearest cluster
	for _, point := range points {
		var minDistance float64 = math.MaxFloat64
		var nearestCluster GeospatialPoint

		for clusterCentroid := range clusters {
			distance := haversineDistance(point, clusterCentroid)
			if distance < minDistance {
				minDistance = distance
				nearestCluster = clusterCentroid
			}
		}

		if minDistance > threshold {
			// If the point is too far from any existing cluster, create a new cluster
			clusters[point] = []GeospatialPoint{point}
		} else {
			// Add the point to the nearest cluster
			clusters[nearestCluster] = append(clusters[nearestCluster], point)
		}
	}

	// Recalculate centroids for each cluster
	for centroid, points := range clusters {
		newCentroid := calculateCentroid(points)
		delete(clusters, centroid)
		clusters[newCentroid] = points
	}

	return clusters
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
		{40.7128, -74.0060},
		{51.5074, -0.1278},
	}

	// Cluster the points
	clusters := clusterCoordinates(points, threshold)

	// Display the clustered geospatial points
	for centroid, groupPoints := range clusters {
		fmt.Printf("Cluster Centroid: (%f, %f)\n", centroid.Latitude, centroid.Longitude)
		for _, point := range groupPoints {
			fmt.Printf("(%f, %f)\n", point.Latitude, point.Longitude)
		}
		fmt.Println()
	}
}
