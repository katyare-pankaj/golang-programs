package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robfig/cron"
)

const (
	dataInterval    = "*/1 * * * *" // Generate data every minute
	maxDataPoints   = 100
	causalThreshold = 0.9
)

type dataPoint struct {
	value     float64
	timestamp time.Time
}

var (
	dataPoints  []dataPoint
	causalModel []float64
)

func generateData() {
	newValue := rand.Float64() * 100
	dataPoints = append(dataPoints, dataPoint{value: newValue, timestamp: time.Now()})
	if len(dataPoints) > maxDataPoints {
		dataPoints = dataPoints[1:]
	}
}

func buildCausalModel() {
	if len(dataPoints) < 2 {
		return
	}

	causalModel = make([]float64, len(dataPoints)-1)
	for i := 0; i < len(causalModel); i++ {
		causalModel[i] = dataPoints[i+1].value / dataPoints[i].value
	}
}

func filterData() {
	if len(causalModel) == 0 {
		return
	}

	for i := 1; i < len(dataPoints); i++ {
		expectedValue := dataPoints[i-1].value * causalModel[i-1]
		if absFloat64(dataPoints[i].value-expectedValue) > causalThreshold {
			// Data point is an outlier, filter it
			dataPoints[i].value = expectedValue
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
	rand.Seed(time.Now().UnixNano())
	c := cron.New()
	c.AddFunc(dataInterval, generateData)
	c.AddFunc(dataInterval, buildCausalModel)
	c.AddFunc(dataInterval, filterData)
	c.Start()

	for {
		time.Sleep(time.Second)
		fmt.Println("Filtered Data:")
		for _, dp := range dataPoints {
			fmt.Printf("%v: %.2f\n", dp.timestamp, dp.value)
		}
		fmt.Println("--------------------")
	}
}
