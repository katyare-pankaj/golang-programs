package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

// Callback function to be called during model training
func trainingCallback(model *trees.DecisionTree, data base.FixedDataGrid, iteration int) {
	// You can add your explainable AI logic here, for example, model evaluation or feature importance calculation.
	// Keep in mind that this function will be called multiple times during training, so keep it efficient.
	fmt.Printf("Iteration %d: Model Accuracy = %.2f\n", iteration, model.Accuracy(data))
}

func main() {
	// Sample dataset for demonstration purposes
	irisData, err := base.ParseCSV("iris.csv")
	if err != nil {
		panic(err)
	}

	// Create a decision tree model
	dt := trees.NewDecisionTree()

	// Set the training callback function
	dt.SetTrainingCallback(trainingCallback)

	// Train the model
	if err := dt.Train(irisData); err != nil {
		panic(err)
	}

	// Make predictions on the test data
	predictions, err := dt.Predict(irisData)
	if err != nil {
		panic(err)
	}

	// Print the accuracy of the model on the test data
	fmt.Println("Model Accuracy:", dt.Accuracy(irisData))

	// Example of explainability: Print the feature importance
	featureImportance := dt.FeatureImportance()
	fmt.Println("Feature Importance:")
	for feature, importance := range featureImportance {
		fmt.Printf("%s: %.2f\n", feature, importance)
	}
}
