package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// PlayerPerformance represents a player's performance in a game
type PlayerPerformance struct {
	Name     string
	Points   float64
	Rebounds float64
	Assists  float64
	Steals   float64
	Blocks   float64
}

func main() {
	// Sample player performance data
	performances := []PlayerPerformance{
		{"Player1", 25, 6, 4, 2, 1},
		{"Player2", 18, 5, 3, 3, 2},
		{"Player3", 22, 4, 6, 1, 0},
		{"Player4", 16, 7, 2, 2, 1},
		{"Player5", 20, 5, 5, 0, 2},
	}

	// Calculate player performance metrics using linear algebra
	evaluatePlayerPerformance(performances)
}

func evaluatePlayerPerformance(performances []PlayerPerformance) {
	// Create a matrix to store player statistics
	numPlayers := len(performances)
	statMatrix := mat.NewDense(numPlayers, 5, nil)
	for i, performance := range performances {
		statMatrix.Set(i, 0, performance.Points)
		statMatrix.Set(i, 1, performance.Rebounds)
		statMatrix.Set(i, 2, performance.Assists)
		statMatrix.Set(i, 3, performance.Steals)
		statMatrix.Set(i, 4, performance.Blocks)
	}

	// Calculate mean and standard deviation for each statistic
	means := calculateMean(statMatrix)
	stds := calculateStdDev(statMatrix, means)

	// Calculate Z-scores for each player
	zScores := mat.NewDense(numPlayers, 5, nil)
	for i := 0; i < numPlayers; i++ {
		for j := 0; j < 5; j++ {
			value := statMatrix.At(i, j)
			zScores.Set(i, j, (value-means.AtVec(j))/stds.AtVec(j))
		}
	}

	// Interpret Z-scores:
	// Z-score > 1: Above average performance
	// Z-score < -1: Below average performance
	// Z-score between -1 and 1: Average performance

	for i, performance := range performances {
		fmt.Printf("Player: %s\n", performance.Name)
		for j := 0; j < 5; j++ {
			statName := []string{"Points", "Rebounds", "Assists", "Steals", "Blocks"}[j]
			zScore := zScores.At(i, j)
			fmt.Printf("%s: %.2f ", statName, zScore)
			if zScore > 1 {
				fmt.Println("(Above average)")
			} else if zScore < -1 {
				fmt.Println("(Below average)")
			} else {
				fmt.Println("(Average)")
			}
		}
		fmt.Println("--------------------")
	}
}

func calculateMean(data *mat.Dense) *mat.VecDense {
	rows, cols := data.Dims()
	means := mat.NewVecDense(cols, data.RawMatrix().Data)

	for j := 0; j < cols; j++ {
		sum := 0.0
		for i := 0; i < rows; i++ {
			sum += data.At(i, j)
		}
		means.SetVec(j, sum/float64(rows))
	}
	return means
}

func calculateStdDev(data *mat.Dense, means *mat.VecDense) *mat.VecDense {
	rows, cols := data.Dims()
	stds := mat.NewVecDense(cols, means.RawVector().Data)

	for j := 0; j < cols; j++ {
		sum := 0.0
		for i := 0; i < rows; i++ {
			diff := data.At(i, j) - means.AtVec(j)
			sum += diff * diff
		}
		variance := sum / float64(rows-1)
		stds.SetVec(j, math.Sqrt(variance))
	}
	return stds
}
