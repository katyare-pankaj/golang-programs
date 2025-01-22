package main

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y float64
}

func crossProduct(o, a, b Point) float64 {
	return (a.X-o.X)*(b.Y-o.Y) - (b.X-o.X)*(a.Y-o.Y)
}

func ConvexHull(points []Point) []Point {
	// Sort the points by increasing x-coordinate, breaking ties by y-coordinate
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X || (points[i].X == points[j].X && points[i].Y < points[j].Y)
	})

	n := len(points)
	hull := make([]Point, 0, n) // Preallocate memory for the hull

	for i := 0; i < n; i++ {
		// Process all points, checking if they make a right turn with the last two points of the hull
		for len(hull) >= 2 && crossProduct(hull[len(hull)-2], hull[len(hull)-1], points[i]) < 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, points[i])
	}

	return hull
}

func main() {
	points := []Point{
		{1, 1}, {3, 3}, {2, 2}, {5, 5}, {4, 1}, {0, 4},
		{1, 1}, // Duplicate point
	}

	fmt.Println("Input points:")
	for _, p := range points {
		fmt.Printf("(%v, %v) ", p.X, p.Y)
	}
	fmt.Println()

	// Remove duplicate points using a map (in order to preserve the original order)
	uniquePoints := map[Point]struct{}{}
	for _, p := range points {
		uniquePoints[p] = struct{}{}
	}

	var sortedPoints []Point
	for p := range uniquePoints {
		sortedPoints = append(sortedPoints, p)
	}

	hull := ConvexHull(sortedPoints)

	fmt.Println("Convex Hull:")
	for _, p := range hull {
		fmt.Printf("(%v, %v) ", p.X, p.Y)
	}
	fmt.Println()
}
