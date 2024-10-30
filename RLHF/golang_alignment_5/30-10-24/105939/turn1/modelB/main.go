package main

import (
	"fmt"
)

const windowSize = 10

func calculateMovingAverage(scores []int) float64 {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}

func predictTrend(playerScores []int) float64 {
	var trends []float64
	for i := 0; i < len(playerScores)-windowSize+1; i++ {
		windowScores := playerScores[i : i+windowSize]
		trends = append(trends, calculateMovingAverage(windowScores))
	}
	return calculateMovingAverage(trends)
}

func main() {
	playerScores := []int{100, 80, 90, 120, 110, 105, 95, 130, 115, 125, 140, 135}
	trend := predictTrend(playerScores)
	fmt.Println("Predicted Trend:", trend)
}
