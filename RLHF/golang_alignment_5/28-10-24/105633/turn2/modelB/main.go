package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

// Player represents a basketball player's statistics
type Player struct {
	Name   string
	Points float64
	Rebounds float64
	Assists float64
	Steals  float64
	Blocks  float64
}

func main() {
	// Dummy data for several players
	players := []Player{
		{"Player1", 20.5, 5.0, 3.0, 2.0, 1.0},
		{"Player2", 18.0, 4.0, 6.0, 3.0, 0.5},
		{"Player3", 15.0, 6.0, 2.0, 1.5, 2.0},
		{"Player4", 22.0, 3.0, 4.0, 2.5, 1.0},
		{"Player5", 19.0, 5.0, 5.0, 1.0, 2.5},
	}

	// Create a matrix to store player statistics
	numPlayers := len(players)
	statMatrix := mat.NewDense(numPlayers, 6, nil)
	for i, player := range players {
		statMatrix.Set(i, 0, player.Points)
		statMatrix.Set(i, 1, player.Rebounds)
		statMatrix.Set(i, 2, player.Assists)
		statMatrix.Set(i, 3, player.Steals)
		statMatrix.Set(i, 4, player.Blocks)
	}

	// Normalize each column of the matrix (optional)
	for j := 0; j < statMatrix.Cols(); j++ {
		column := statMatrix.ColView(j)
		mean := stat.Mean(column, nil)
		stdDev := stat.StdDev(column, nil)
		for i := 0; i < statMatrix.Rows(); i++ {
			statMatrix.Set(i, j, (statMatrix.At(i, j)-mean)/stdDev)
		}
	}

	// Define weights for each performance metric (optional)
	weights := mat.NewVecDense(6, []float64{0.2, 0.2, 0.2, 0.15, 0.15})

	// Calculate weighted linear combination scores
	weightedScores := mat.NewVecDense(numPlayers, nil)
	for i := 0; i < numPlayers; i++ {
		row := statMatrix.RowView(i)
		weightedScore := mat.Dot(row, weights)
		weightedScores.SetVec(i, weightedScore)
	}

	// Perform PCA to reduce dimensionality (optional)
	pca := NewPCA(statMatrix)
	reducedDimension := 2 // Choose the desired number of reduced dimensions