package main  
import (  
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	dataInterval = "*/1 * * * *" // Generate data every minute
	maxDataPoints = 100
	causalThreshold = 0.05
)

type dataPoint struct {
	heartRate   float64
	breathingRate float64
	timestamp    time.Time
}

var (
	dataPoints []dataPoint
	causalModel struct {
		heartRateToBreathing float64
	}
)

func generateData() {
	newHeartRate := rand.Float64() * 180 + 40 // Normal heart rate range for adults
	newBreathingRate := rand.Float64() * 30 + 10
	dataPoints = append(dataPoints, dataPoint{heartRate: newHeartRate, breathingRate: newBreathingRate, timestamp: time.Now()})
	if len(dataPoints) > maxDataPoints {
		dataPoints = dataPoints[1:]
	}
}

func buildCausalModel() {
	if len(dataPoints) < 2 {
		return
	}

	// Simple linear model: Heart rate = alpha * Breathing rate + beta
	var sumXY, sumX, sumY float64
	for _, dp := range dataPoints {
		sumXY += dp.heartRate * dp.breathingRate
		sumX += dp.breathingRate
		sumY += dp.heartRate
	}

	n := float64(len(dataPoints))
	causalModel.heartRateToBreathing = (n*sumXY - sumX*sumY) / (n*sumX*sumX - sumX*sumX)
}

func filterData() {
	if causalModel.heartRateToBreathing == 0 {
		return
	}

	for i := 1; i < len(dataPoints); i++ {
		expectedHeartRate := causalModel.heartRateToBreathing * dataPoints[i].breathingRate
		if absFloat64(dataPoints[i].heartRate-expectedHeartRate) > causalThreshold {
			// Data point is an outlier, filter it
			dataPoints[i].heartRate = expectedHeartRate
		}
	}
}

func absFloat64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {