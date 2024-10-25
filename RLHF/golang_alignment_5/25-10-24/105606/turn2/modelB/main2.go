package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Sample log data with features (e.g., request count, response time) and labels (success/failure)
	data := [][]string{
		{"100", "120", "success"},
		{"80", "100", "success"},
		{"150", "160", "failure"},
		{"110", "130", "success"},
		{"130", "140", "success"},
		{"140", "180", "failure"},
		{"160", "110", "success"},
		{"180", "130", "failure"},
		{"170", "120", "success"},
		{"200", "140", "failure"},
	}
	// Instantiate a new KNN classifier
	knn := knn.NewKNN(3)
	// Prepare the dataset
	dataset := base.LoadLinesDataset(data)
	// Split the data into training and test sets
	trainingSet, _ := base.InstancesTrainTestSplit(dataset, 0.8)
	// Train the KNN model
	err := knn.Fit(trainingSet)
	if err != nil {
		fmt.Println("Error training KNN model:", err)
		return
	}

}
