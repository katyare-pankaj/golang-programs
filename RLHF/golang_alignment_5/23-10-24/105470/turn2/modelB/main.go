package mlpdsl

import (
	"math/rand"
	"testing"
)

// MLP stands for Multi-Layer Perceptron
type MLP struct {
	layers         []int         // Number of neurons in each layer
	weights        [][][]float64 // Weights between neurons in each layer
	biases         [][]float64   // Biases for neurons in each layer
	activationFunc ActivationFunc
}

// ActivationFunc defines the activation function used for neurons
type ActivationFunc func(float64) float64

func TestNewMLP(t *testing.T) {
	// Assert MLP structure is created correctly
	mlp := NewMLP(2, 3, 1, ReLU)
	if len(mlp.layers) != 3 {
		t.Errorf("Expected 3 layers, got %d", len(mlp.layers))
	}
	// Add more assertions for weights, biases, etc.
}

func BenchmarkFeedForward(b *testing.B) {
	// Measure the time taken to perform feed-forward for a single input
	mlp := NewMLP(100, 50, 10, ReLU)
	input := make([]float64, 100)
	for i := range input {
		input[i] = rand.Float64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mlp.FeedForward(input)
	}
}

// Stochastic Gradient Descent (SGD) implementation
func (mlp *MLP) SGD(trainingData [][]float64, targets []float64, learningRate float64, epochs int, batchSize int) {
	for epoch := 0; epoch < epochs; epoch++ {
		for i := 0; i < len(trainingData); i += batchSize {
			batchInputs := trainingData[i:min(i+batchSize, len(trainingData))]
			batchTargets := targets[i:min(i+batchSize, len(targets))]
			mlp.trainBatch(batchInputs, batchTargets, learningRate)
		}
	}
}

func (mlp *MLP) trainBatch(inputs [][]float64, targets []float64, learningRate float64) {
	// Implement the SGD training logic for a batch
	// ... (details omitted)
}

func TestSGD(t *testing.T) {
	// Assert that SGD trains the MLP with reasonable accuracy
	mlp := NewMLP(2, 3, 1, ReLU)
	trainingData := [][]float64{{1, 1}, {0, 1}, {1, 0}, {0, 0}}
	targets := []float64{1, 1, 0, 0}
	mlp.SGD(trainingData, targets, 0.01, 100, 1)
	// Perform additional assertions to evaluate the accuracy
	// ...
}
func TestSGD(t *testing.T) {
	// Assert that SGD trains the MLP with reasonable accuracy
	mlp := NewMLP(2, 3, 1, ReLU)
	trainingData := [][]float64{{1, 1}, {0, 1}, {1, 0}, {0, 0}}
	targets := []float64{1, 1, 0, 0}
	mlp.SGD(trainingData, targets, 0.01, 100, 1)
	// Perform additional assertions to evaluate the accuracy
	// ...
}
