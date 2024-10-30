package main

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CalculateStatistic", func() {
	Describe("Calculation", func() {
		It("should calculate the sum correctly", func() {
			data := []int{1, 2, 3, 4}
			result := CalculateStatistic(data)
			Expect(result).To(Equal(10))
		})
		It("should handle larger data sets correctly", func() {
			data := make([]int, 1000)
			for i := 0; i < len(data); i++ {
				data[i] = i
			}
			result := CalculateStatistic(data)
			Expect(result).To(Equal(499500))
		})
	})
	Describe("Parallelization", func() {
		It("should perform calculations in parallel", func() {
			data := []int{1, 2, 3, 4, 5}
			startTime := time.Now()
			CalculateStatistic(data)
			duration := time.Since(startTime)
			// Assert that the execution time is reasonable for parallel processing
			Expect(duration).Should(BeNumerically("<", 10*time.Millisecond))
		})
	})
})
