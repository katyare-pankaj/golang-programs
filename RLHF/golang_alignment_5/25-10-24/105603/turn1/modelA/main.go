package main

import (
	"fmt"

	"github.com/sugarme/gotch"
	"github.com/sugarme/gotch/nn"
)

func main() {
	// Initialize Gotch (NumGO)
	gotch.Init()
	defer gotch.Close()

	// Create a simple linear regression model
	model := nn.Sequential(
		nn.Linear(2, 1),
	)

	// Sample input data and labels
	inputs := gotch.NewTensor([][]float32{{1, 2}, {3, 4}, {5, 6}}, gotch.Float32)
	labels := gotch.NewTensor([][]float32{{3}, {5}, {7}}, gotch.Float32)

	// Optimize the model using SGD
	optimizer := gotch.NewSGD(model.Parameters(), 0.01, 0.9)

	// Training loop
	for epoch := 0; epoch < 100; epoch++ {
		// Forward pass
		predictions := model.Forward(inputs)
		loss := nn.MSELoss(predictions, labels)

		// Backward pass and optimize
		loss.Backward()
		optimizer.Step()
		optimizer.ZeroGrad()

		// Print loss
		if epoch%10 == 0 {
			fmt.Printf("Epoch %d, Loss: %.6f\n", epoch, loss.Item())
		}
	}
}
