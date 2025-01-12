package main  
import (  
    "fmt"
    "math"
    "math/rand"
    "time"
)

// Define the GeospatialPoint struct
type GeospatialPoint struct {  
    Latitude  float64
    Longitude float64
}  

// Function to calculate the euclidean distance between two points
func euclideanDistance(point1, point2 GeospatialPoint) float64 {  
    dx := point2.Latitude - point1.Latitude
    dy := point2.Longitude - point1.Longitude
    return math.Sqrt(dx*dx + dy*dy)
}

// K-means clustering algorithm
func kmeansClustering(points []GeospatialPoint, k int) [][]GeospatialPoint {  
    rand.Seed(time.Now().UnixNano())

    // Initialize k random centroids as clusters
    centroids := make([]GeospatialPoint, k)
    for i := 0; i < k; i++ {
        centroidIndex := rand.Intn(len(points))
        centroids[i] = points[centroidIndex]
    }

    maxIterations := 100
    for i := 0; i < maxIterations; i++ {  
        clusters := make([][]GeospatialPoint, k)
        for _, point := range points {  
            // Find the closest centroid for each point
            minDistance := math.MaxFloat64
            closestCluster := -1
            for j, centroid := range centroids {  
                distance := euclideanDistance(point, centroid)
                if distance < minDistance {  
                    minDistance = distance
                    closestCluster = j
                }
            }

            // Add the point to its closest cluster
            clusters[closestCluster] = append(clusters[closestCluster], point)
        }

        // Update centroids as the mean of the points in each cluster
        newCentroids := make([]GeospatialPoint, k)
        for j, cluster := range clusters {  
            var sumLat, sumLon float64
            if len(cluster) > 0 {
                for _, point := range cluster {  
                    sumLat += point.Latitude
                    sumLon += point.Longitude
                }
                newCentroids[j].Latitude = sumLat / float64(len(cluster))
                newCentroids[j].Longitude = sumLon / float64(len(cluster))
            } else {
                // If a cluster is empty, choose a random point as the new centroid
                newCentroidIndex := rand.Intn(len(points))
                newCentroids[j] = points[newCentroidIndex]
            }
        }

        // Check for convergence
        if converged(centroids, newCentroids) {  
            return clusters
        }
        centroids = newCentroids
    }

    // If max iterations are reached, return the clusters
    return clusters
}

// Function to check for convergence
func converged(centroids1, centroids2 []GeospatialPoint) bool {  
    const convergenceThreshold = 0.0001

    for i := 0; i < len(centroids1); i++ {  
        distance := euclideanDistance(centroids1[i], centroids2[i])
        if distance > convergenceThreshold {  
            return false
        }
    }
    return true
}
  
func main() {  
    const k = 3 // Number of clusters

    // Define a slice of geospatial points
    points := []GeospatialPoint{
        {32.6514, -161.4333},
        {37.7747, -122.4182},  
        {31.2345, -101.5361},  
        {38.8977, -77.0365},