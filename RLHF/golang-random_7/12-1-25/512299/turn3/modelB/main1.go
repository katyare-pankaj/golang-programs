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

// Function to calculate the Mahalanobis distance between two geospatial points
func mahalanobisDistance(point1, point2 GeospatialPoint, covariance [2][2]float64) float64 {
	diffLat := point2.Latitude - point1.Latitude
	diffLon := point2.Longitude - point1.Longitude
	invCov := [2][2]float64{
		{1 / covariance[0][0], -covariance[0][1] / (covariance[0][0] * covariance[1][1])},
		{-covariance[1][0] / (covariance[0][0] * covariance[1][1]), 1 / covariance[1][1]},
	}
	distance := math.Sqrt(diffLat*invCov[0][0]*diffLat + diffLon*invCov[0][1]*diffLat + diffLat*invCov[1][0]*diffLon + diffLon*invCov[1][1]*diffLon)
	return distance
}

// Function to calculate the covariance matrix of geospatial points
func calculateCovariance(points []GeospatialPoint) [2][2]float64 {
	n := float64(len(points))
	if n == 0 {
		return [2][2]float64{} // Return empty matrix if no points
	}

	sumLat := 0.0
	sumLon := 0.0
	for _, point := range points {
		sumLat += point.Latitude
		sumLon += point.Longitude
	}
	meanLat := sumLat / n
	meanLon := sumLon / n

	var covariance [2][2]float64
	for _, point := range points {
		diffLat := point.Latitude - meanLat
		diffLon := point.Longitude - meanLon
		covariance[0][0] += diffLat * diffLat
		covariance[0][1] += diffLat * diffLon
		covariance[1][0] += diffLon * diffLat
		covariance[1][1] += diffLon * diffLon
	}
	covariance[0][0] /= n
	covariance[0][1] /= n
	covariance[1][0] /= n
	covariance[1][1] /= n

	return covariance
}

// Function to detect and exclude outliers using Mahalanobis distance
func removeOutliers(points []GeospatialPoint, threshold float64) []GeospatialPoint {
	covariance := calculateCovariance(points)
	meanPoint := GeospatialPoint{Latitude: 0, Longitude: 0}
	for _, point := range points {
		meanPoint.Latitude += point.Latitude
		meanPoint.Longitude += point.Longitude
	}
	meanPoint.Latitude /= float64(len(points))
	meanPoint.Longitude /= float64(len(points))

	var cleanedPoints []GeospatialPoint
	for _, point := range points {
		distance := mahalanobisDistance(point, meanPoint, covariance)
		if distance <= threshold {
			cleanedPoints = append(cleanedPoints, point) // Keep points within the threshold
		}
	}
	return cleanedPoints
}

func main() {
	// Sample geospatial data
	points := []GeospatialPoint{
		{Latitude: 40.7128, Longitude: -74.0060},  // New York
		{Latitude: 34.0522, Longitude: -118.2437}, // Los Angeles
		{Latitude: 41.8781, Longitude: -87.6298},  // Chicago
		{Latitude: 51.5074, Longitude: -0.1278},   // London
		{Latitude: 100.0, Longitude: 200.0},       // Outlier point
	}

	// Set threshold for outlier removal (e.g., distance > 3 standard deviations is an outlier)
	threshold := 3.0

	// Remove outliers
	cleanedPoints := removeOutliers(points, threshold)

	// Print cleaned geospatial points
	fmt.Println("Cleaned Points:")
	for _, point := range cleanedPoints {
		fmt.Printf("Latitude: %.4f, Longitude: %.4f\n", point.Latitude, point.Longitude)
	}
}
