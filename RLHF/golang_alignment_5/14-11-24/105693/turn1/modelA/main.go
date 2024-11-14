package main

const (
	N = 1024 // Matrix size for better performance observation
)

func doWorkWithoutOptimization(a [][]float64) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] *= 2.0 // Simple work on each element
		}
	}
}

func doWorkWithInlining(a [][]float64) {
	const two float64 = 2.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] *= two // Compiler can inline the constant
		}
	}
}

func main() {
	var a [][]float64 = make([][]float64, N, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = float64(i*N + j)
		}
	}

	// Perform work without optimization
	doWorkWithoutOptimization(a)
	// Perform work with inlining optimization
	doWorkWithInlining(a)

	// You can add benchmarking here to measure performance
}
