package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/dereklstinson/gocuda"
	"github.com/dereklstinson/gocuda/gpu"
)

const (
	arraySize = 1024 * 1024
	threshold = 0.5
)

func main() {
	// Initialize CUDA
	dev, err := gocuda.InitDevice(0)
	if err != nil {
		log.Fatalf("Failed to initialize CUDA: %v", err)
	}
	defer dev.Close()

	// Create host arrays
	hA := make([]float32, arraySize)
	hB := make([]float32, arraySize)
	hC := make([]float32, arraySize)

	// Initialize data on the host
	for i := 0; i < arraySize; i++ {
		hA[i] = rand.Float32()
		hB[i] = rand.Float32()
	}

	// Allocate device arrays
	dA, err := gpu.NewArray(hA)
	if err != nil {
		log.Fatalf("Failed to allocate device array A: %v", err)
	}
	defer dA.Free()

	dB, err := gpu.NewArray(hB)
	if err != nil {
		log.Fatalf("Failed to allocate device array B: %v", err)
	}
	defer dB.Free()

	dC, err := gpu.NewArray(hC)
	if err != nil {
		log.Fatalf("Failed to allocate device array C: %v", err)
	}
	defer dC.Free()

	// Launch the kernel
	kernel := `
	__global__ void elementWiseAdd(float* A, float* B, float* C, const float threshold) {
		int idx = blockIdx.x * blockDim.x + threadIdx.x;

		if (idx < {{.ArraySize}}) {
			float a = A[idx];
			float b = B[idx];
			float sum = a + b;
			
			// Branch prediction should work well here with a predictable threshold
			if (sum > threshold) {
				C[idx] = sum;
			} else {
				C[idx] = 0;
			}
		}
	}
	`

	// Create the kernel
	tmpl, err := dev.NewTemplateFromSource(kernel, map[string]interface{}{
		"ArraySize": arraySize,
	})
	if err != nil {
		log.Fatalf("Failed to create kernel template: %v", err)
	}
	defer tmpl.Free()

	kernelConfig := gpu.KernelConfig{
		BlockDim:    gpu.Dim3{arraySize / 1024, 1, 1}, // Adjust block size as needed
		GridDim:     gpu.Dim3{1, 1, 1},
		SharedMem:   0,
		Stream:      nil,
		ExtraParams: []interface{}{dA.Ptr(), dB.Ptr(), dC.Ptr(), threshold},
	}

	err = tmpl.LaunchKernel(kernelConfig)
	if err != nil {
		log.Fatalf("Failed to launch kernel: %v", err)
	}

	// Copy result back to host
	err = dC.CopyToHost(hC)
	if err != nil {
		log.Fatalf("Failed to copy result back to host: %v", err)
	}

	// Validate the result
	correct := true
	for i := 0; i < arraySize; i++ {
		sum := hA[i] + hB[i]
		if sum > threshold && hC[i] != sum {
			correct = false
			break
		} else if sum <= threshold && hC[i] != 0 {
			correct = false
			break
		}
	}

	if correct {
		fmt.Println("Result is correct!")
	} else {
		fmt.Println("Result verification failed!")
	}
}
