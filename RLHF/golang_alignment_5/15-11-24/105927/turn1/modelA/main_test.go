package main

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Polygon struct {
	Coords []coord.Point
}

func rasterizePolygon(polygon Polygon, cellSize float64) (rasterArea float64) {
	// Implementation of rasterization algorithm here
	// For simplicity, let's assume a basic bounding box approach
	minX, minY := float64(1000), float64(1000)
	maxX, maxY := float64(0), float64(0)

	for _, c := range polygon.Coords {
		minX = math.Min(minX, c.X)
		minY = math.Min(minY, c.Y)
		maxX = math.Max(maxX, c.X)
		maxY = math.Max(maxY, c.Y)
	}

	numCellsX := int((maxX-minX)/cellSize) + 1
	numCellsY := int((maxY-minY)/cellSize) + 1

	rasterArea = float64(numCellsX) * float64(numCellsY) * cellSize * cellSize

	return
}

var _ = Describe("Rasterization", func() {
	Context("when rasterizing a polygon", func() {
		var (
			polygon      Polygon
			cellSize     float64 = 0.1
			expectedArea float64 = 12.5 // Example expected area
			tolerance    float64 = 0.05
		)

		BeforeEach(func() {
			// Example polygon coordinates
			polygon.Coords = []coord.Point{
				{X: 1.0, Y: 1.0}, {X: 5.0, Y: 1.0}, {X: 5.0, Y: 5.0}, {X: 1.0, Y: 5.0}, {X: 1.0, Y: 1.0},
			}
		})

		It("should approximate the polygon's area", func() {
			rasterizedArea := rasterizePolygon(polygon, cellSize)

			Expect(rasterizedArea).To(BeNumerically("~=", expectedArea, tolerance))
		})
	})
})
