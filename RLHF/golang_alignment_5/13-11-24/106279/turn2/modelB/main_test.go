package main

import "testing"

func benchmarkCelsiusToFahrenheitReadable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		celsiusToFahrenheitReadable(25.0)
	}
}

func benchmarkCelsiusToFahrenheitEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		celsiusToFahrenheitEfficient(float32(25.0))
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	// Running the benchmark for celsiusToFahrenheitReadable
	t.Run("BenchmarkCelsiusToFahrenheitReadable", func(t *testing.T) {
		res := testing.Benchmark(benchmarkCelsiusToFahrenheitReadable)
		if res.N == 0 {
			t.Error("BenchmarkCelsiusToFahrenheitReadable did not run")
		}
	})

	// Running the benchmark for celsiusToFahrenheitEfficient
	t.Run("BenchmarkCelsiusToFahrenheitEfficient", func(t *testing.T) {
		res := testing.Benchmark(benchmarkCelsiusToFahrenheitEfficient)
		if res.N == 0 {
			t.Error("BenchmarkCelsiusToFahrenheitEfficient did not run")
		}
	})
}
