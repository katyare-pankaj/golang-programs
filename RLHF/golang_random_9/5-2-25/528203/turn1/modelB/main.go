package main

import (
	"fmt"
	"os/exec"
	"time"
)

func primeSieve(n int) []int {
	primes := make([]int, 0, n/10)
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return primes
}
func main() {
	n := 10000000 // Example large number
	start := time.Now()
	primes := primeSieve(n)
	elapsed := time.Since(start)
	fmt.Printf("Found %d primes in %.2fs\n", len(primes), elapsed.Seconds())

	// Running go vet
	err := runGoVet()
	if err != nil {
		fmt.Printf("Error running go vet: %v\n", err)
	}

	// Running golint
	err = runGoLint()
	if err != nil {
		fmt.Printf("Error running go lint: %v\n", err)
	}
}

// Function to run go vet command
func runGoVet() error {
	cmd := exec.Command("go", "vet", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go vet failed: %v\nOutput:\n%s", err, output)
	}
	fmt.Printf("go vet output:\n%s", output)
	return nil
}

// Function to run golint command
func runGoLint() error {
	cmd := exec.Command("golint", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("golint failed: %v\nOutput:\n%s", err, output)
	}
	fmt.Printf("golint output:\n%s", output)
	return nil
}
