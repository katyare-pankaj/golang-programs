package main

import (
	"fmt"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// Player represents a basketball player's statistics
type Player struct {
	Name     string
	Points   float64
	Rebounds float64
	Assists  float64
}

func main() {
	// Dummy data for several players
	players := []Player{
		{"Player1", 20.5, 5.0, 3.0},
		{"Player2", 18.0, 4.0, 6.0},
		{"Player3", 15.0, 6.0, 2.0},
		{"Player4", 22.0, 3.0, 4.0},
		{"Player5", 19.0, 5.0, 5.0},
	}

	// Create a matrix to store player statistics
	numPlayers := len(players)
	statMatrix := mat.NewDense(numPlayers, 3, nil)
	for i, player := range players {
		statMatrix.Set(i, 0, player.Points)
		statMatrix.Set(i, 1, player.Rebounds)
		statMatrix.Set(i, 2, player.Assists)
	}

	// Perform K-means clustering to group similar players
	k := 2 // Number of clusters
	centroids := initializeCentroids(statMatrix, k)
	labels := kmeansClustering(statMatrix, centroids, k)

	// Display cluster results
	fmt.Println("Cluster Results:")
	for i, label := range labels {
		fmt.Printf("Player %s: Cluster %d\n", players[i].Name, label)
	}

	// Optimize team composition for a given scoring ratio
	scoringRatio := 0.6 // Desired scoring ratio (points/rebounds)
	optimalLineup := optimizeTeamComposition(statMatrix, labels, scoringRatio)

	fmt.Println("\nOptimal Team Composition:")
	for _, i := range optimalLineup {
		fmt.Printf("Player %s\n", players[i].Name)
	}
}

// initializeCentroids randomly selects k centroids from the data
func initializeCentroids(data *mat.Dense, k int) *mat.Dense {
	rows, cols := data.Dims()
	centroids := mat.NewDense(k, cols, nil)
	for i := 0; i < k; i++ {
		row := rand.Intn(rows)
		for j := 0; j < cols; j++ {
			centroids.Set(i, j, data.At(row, j))
		}
	}
	return centroids
}

// kmeansClustering performs K-means clustering on the given data
func kmeansClustering(data *mat.Dense, centroids *mat.Dense, k int) []int {
	rows, _ := data.Dims()
	labels := make([]int, rows)

	for iter := 0; iter < 10; iter++ { // Maximum iterations
		// Assign data points to clusters
		for i := 0; i < rows; i++ {
			minDist := float64(math.MaxFloat64)
			var minLabel int
			for j := 0; j < k; j++ {
				dist := euclideanDistance(data.RowView(i), centroids.RowView(j))
				if dist < minDist {
					minDist = dist
					minLabel = j
				}
			}
			labels[i] = minLabel
		}

		// Recalculate centroids
		newCentroids := mat.NewDense(k, data.Cols(), nil)
		count := make([]int, k)
		for i := 0; i < rows; i++ {
			label := labels[i]
			for j := 0; j < data.Cols(); j++ {
				newCentroids.Set(label, j, newCentroids.At(label, j)+data.At(i, j))
			}
			count[label]++
		}
		for i := 0; i < k; i++ {
			if count[i] > 0 {
				for j := 0; j < data.Cols(); j++ {
					newCentroids.Set(i, j, newCentroids.At(i, j)/float64(count[i]))
				}
			}
		}

		// Check for convergence
		if mat.Equal(centroids, newCentroids) {
			break
		}
		centroids = newCentroids
	}

	return labels
}

func euclideanDistance(row1, row2 mat.Vector) float64 {
	sum := 0.0
	for i := 0; i < row1.Len(); i++ {
		diff := row1.AtVec(i) - row2.AtVec(i)
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func optimizeTeamComposition(data *mat.Dense, labels []int, scoringRatio float64) []int {
	rows, cols := data.Dims()
	// Create a matrix to store player statistics for each cluster
	clusterStats := make([][]float64, 2) // Assuming 2 clusters for simplicity
	for i := 0; i < len(clusterStats); i++ {
		clusterStats[i] = make([]float64, cols)
	}

	for i := 0; i < rows; i++ {
		label := labels[i]
		for j := 0; j < cols; j++ {
			clusterStats[label][j] += data.At(i, j)
		}
	}

	// Perform linear programming to optimize the lineup
	// (Implementation details omitted for brevity)
	// For this example, we'll just pick the players with the highest scoring ratio
	var optimalLineup []int
	for label := 0; label < len(clusterStats); label++ {
		bestPlayerIndex := -1
		maxRatio := -1.0
		for i := 0; i < rows; i++ {
			if labels[i] == label {
				points := data.At(i, 0)
				rebounds := data.At(i, 1)
				ratio := points / rebounds
				if ratio > maxRatio {
					maxRatio = ratio
					bestPlayerIndex = i
				}
			}
		}
		optimalLineup = append(optimalLineup, bestPlayerIndex)
	}

	return optimalLineup
}
