package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hybridgroup/gocuda"
)

const (
	vectorSize = 1024 * 1024
)

func main() {
	// Initialize CUDA
	err := gocuda.Init()
	if err != nil {
		panic(err)
	}
	defer gocuda.Shutdown()

	// Create device vector
	devVectorA, err := gocuda.NewMalloc(vectorSize * 4)
	if err != nil {
		panic(err)
	}
	defer devVectorA.Free()

	devVectorB, err := gocuda.NewMalloc(vectorSize * 4)
	if err != nil {
		panic(err)
	}
	defer devVectorB.Free()

	devVectorC, err := gocuda.NewMalloc(vectorSize * 4)
	if err != nil {
		panic(err)
	}
	defer devVectorC.Free()

	// Generate random data for vectors A and B
	rand.Seed(time.Now().UnixNano())
	vectorA := make([]float32, vectorSize)
	vectorB := make([]float32, vectorSize)
	for i := 0; i < vectorSize; i++ {
		vectorA[i] = rand.Float32()
		vectorB[i] = rand.Float32()
	}

	// Copy data to device
	err = gocuda.MemcpyHtoD(devVectorA, vectorA)
	if err != nil {
		panic(err)
	}
	err = gocuda.MemcpyHtoD(devVectorB, vectorB)
	if err != nil {
		panic(err)
	}

	// Launch CUDA kernel to perform vector addition
	kernel := `
        __global__ void vectorAdd(float* a, float* b, float* c) {
                int i = threadIdx.x + blockIdx.x * blockDim.x;

                if (i < 1024 * 1024) {
                        c[i] = a[i] + b[i];
                }
        }
        `
	module, err := gocuda.ParseModule(kernel)
	if err != nil {
		panic(err)
	}
	defer module.Free()

	vectorAddFunc, err := module.GetFunction("vectorAdd")
	if err != nil {
		panic(err)
	}

	gridDimX := (vectorSize + 255) / 256
	blockDimX := 256
	gridDim := gocuda.Dim3{gridDimX, 1, 1}
	blockDim := gocuda.Dim3{blockDimX, 1, 1}

	err = vectorAddFunc.Launch(gridDim, blockDim, devVectorA, devVectorB, devVectorC)
	if err != nil {
		panic(err)
	}

	// Copy result from device to host
	result := make([]float32, vectorSize)
	err = gocuda.MemcpyDtoH(result, devVectorC)
	if err != nil {
		panic(err)
	}

	// Verify the result
	for i := 0; i < vectorSize; i++ {
		if result[i] != vectorA[i]+vectorB[i] {
			fmt.Printf("Error: Result mismatch at index %d\n", i)
			return
		}
	}

	fmt.Println("Vector addition completed successfully")
}
