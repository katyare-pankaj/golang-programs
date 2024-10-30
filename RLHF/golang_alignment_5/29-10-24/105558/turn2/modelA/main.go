package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	maxDataPoints   = 100
	causalThreshold = 0.1
)

type dataPoint struct {
	sensorID  string
	value     float64
	timestamp time.Time
}

var (
	dataPoints  []dataPoint
	causalModel map[string]map[string]float64
)

func generateData() {
	sensors := []string{"heartRate", "accelerometerX", "accelerometerY", "accelerometerZ"}
	for _, sensor := range sensors {
		newValue := generateSensorData(sensor)
		dataPoints = append(dataPoints, dataPoint{sensorID: sensor, value: newValue, timestamp: time.Now()})
		if len(dataPoints) > maxDataPoints {
			dataPoints = dataPoints[1:]
		}
	}
}

func generateSensorData(sensorID string) float64 {
	// Simulate data generation with noise for different sensors
	baseValue := 0.0
	noise := 0.0
	switch sensorID {
	case "heartRate":
		baseValue = 60 + rand.Float64()*20
		noise = 1.0 + rand.Float64()*2
	case "accelerometerX", "accelerometerY", "accelerometerZ":
		baseValue = rand.Float64()*2 - 1
		noise = 0.05 + rand.Float64()*0.1
	}
	return baseValue + noise*rand.NormFloat64()
}

func buildCausalModel() {
	causalModel = make(map[string]map[string]float64)

	// Simplified causal relationships for demonstration purposes
	causalModel["heartRate"] = map[string]float64{
		"accelerometerX": 0.5,
		"accelerometerY": 0.5,
		"accelerometerZ": 0.5,
	}

	for _, sensor := range []string{"accelerometerX", "accelerometerY", "accelerometerZ"} {
		causalModel[sensor] = map[string]float64{}
	}
}

func filterData() {
	if causalModel == nil {
		return
	}

	for i := 1; i < len(dataPoints); i++ {
		currentPoint := dataPoints[i]
		expectedValue := predictExpectedValue(currentPoint)
		if math.Abs(currentPoint.value-expectedValue) > causalThreshold {
			// Data point is an outlier, filter it
			dataPoints[i].value = expectedValue
		}
	}
}

func predictExpectedValue(point dataPoint) float64 {
	expectedValue := 0.0
	influencingSensors := causalModel[point.sensorID]
	for sensorID, causalFactor := range influencingSensors {
		for _, dp := range dataPoints {
			if dp.sensorID == sensorID {
				expectedValue += dp.value * causalFactor
				break
			}
		}
	}
	return expectedValue
}

func main() {
	rand.Seed(time.Now().UnixNano())

	buildCausalModel()

	for i := 0; i < 100; i++ {
		generateData()
		filterData()
		fmt.Println("Filtered Data:")
		for _, dp := range dataPoints {
			fmt.Printf("%s: %.2f\n", dp.sensorID, dp.value)
		}
		fmt.Println("--------------------")
		time.Sleep(time.Second)
	}
}
