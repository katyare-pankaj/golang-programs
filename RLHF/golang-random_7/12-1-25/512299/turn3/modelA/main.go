package main

import (
	"fmt"
	"math"
	"sort"
)

// Define the GeospatialPoint struct
type GeospatialPoint struct {
	Latitude  float64
	Longitude float64
	Weight    float64 // Weight of the point
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

// Function to recalculate centroids with weights
func calculateCentroidWithWeights(points []GeospatialPoint) GeospatialPoint {
	n := float64(len(points))
	if n == 0 {
		return GeospatialPoint{} // Return empty point if no points
	}

	sumWeight := 0.0
	sumLat := 0.0
	sumLon := 0.0

	for _, point := range points {
		sumWeight += point.Weight
		sumLat += point.Latitude * point.Weight
		sumLon += point.Longitude * point.Weight
	}

	return GeospatialPoint{
		Latitude:  sumLat / sumWeight,
		Longitude: sumLon / sumWeight,
		Weight:    1.0, // Centroid's weight is always 1
	}
}

// Function to compute the weighted sum of distances
func weightedDistanceSum(centroids map[GeospatialPoint]int, points []GeospatialPoint) float64 {
	var distanceSum float64

	for centroid, count := range centroids {
		for _, point := range points {
			distance := haversineDistance(centroid, point)
			distanceSum += distance * point.Weight
		}
	}

	return distanceSum
}

// Function to assign points to clusters based on distance
func assignPointsToClusters(centroids map[GeospatialPoint]int, points []GeospatialPoint) map[GeospatialPoint][]GeospatialPoint {
	clusters := make(map[GeospatialPoint][]GeospatialPoint)

	for _, point := range points {
		var minDistance float64 = math.MaxFloat64
		var nearestCluster GeospatialPoint

		for centroid := range centroids {
			distance := haversineDistance(centroid, point)
			if distance < minDistance {
				minDistance = distance
				nearestCluster = centroid
			}
		}

		clusters[nearestCluster] = append(clusters[nearestCluster], point)
	}

	return clusters
}

// Function to perform weighted k-means clustering
func weightedKMeans(points []GeospatialPoint, k int, threshold float64, maxIter int) map[GeospatialPoint][]GeospatialPoint {
	// Initialize centroids with random points
	centroids := make(map[GeospatialPoint]int)
	n := len(points)
	if n < k {
		panic("The number of points must be greater than or equal to the number of clusters.")
	}

	shuffledIndices := sort.Ints(make([]int, n))
	sort.Shuffle(shuffledIndices, func(i, j int) {})

	for i := 0; i < k; i++ {
		centroids[points[shuffledIndices[i]]] = i + 1
	}

	for iter := 0; iter < maxIter; iter++ {
		// Assign points to the closest centroid
		clusters := assignPointsToClusters(centroids, points)

		// Recalculate centroids with weights
		newCentroids := make(map[GeospatialPoint]int)
		for centroid, groupPoints := range clusters {
			newCentroid := calculateCentroidWithWeights(groupPoints)
			newCentroids[newCentroid] = centroids[centroid]
		}

		// Check for convergence
		distanceSum := weightedDistanceSum(newCentroids, points)
		fmt.Printf("Iteration %d: Distance sum = %.2f\n", iter+1, distanceSum)

		if distanceSum < threshold {
			fmt.Printf("Clustering converged after %d iterations\n", iter+1)
			return clusters
		}

		centroids = newCentroids
	}

	fmt.Println("Clustering did not converge after maximum iterations.")
	return clusters
}

func main() {
	const threshold = 1.0 // Proximity threshold for convergence
	const maxIter = 100   // Maximum number of iterations
	const k = 3           // Number of clusters

	// Define a slice of geospatial points
	points := []GeospatialPoint{
		{32.6514, -161.4333, 1.0},
		{37.7747, -122.4182, 1.0},
		{31.2345, -101.5361, 1.0},
		{37.7747, -122.4172, 1.0}, // A point close to the third one
		{38.8977, -77.0365, 1.0},
		{32.6514, -161.4343, 1.0}, // Another point close to the first one
		{40.7128, -74.0060, 0.1},  // An outlier
		{51.5074, -0.1278, 0.1},   // Another outlier
	}

	// Perform weighted k-means clustering
	clusters := weightedKMeans(points, k, threshold, maxIter)

	// Display the clustered geospatial points
	for centroid, groupPoints := range clusters {
		fmt.Printf("Cluster Centroid: (%f, %f)\n", centroid.Latitude, centroid.Longitude)
		for _, point := range groupPoints {
			fmt.Printf("(%f, %f, %f)\n", point.Latitude, point.Longitude, point.Weight)
		}
		fmt.Println()
	}
}
