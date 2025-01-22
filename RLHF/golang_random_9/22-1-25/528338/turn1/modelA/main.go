package main

import (
	"fmt"
	"sort"
)

// Point represents a point in 2D.
type Point struct {
	X, Y int
}

// Cross product of two vectors OA and OB
// A − B
func cross(o, a, b Point) int {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

// ConvexHull computes the convex hull of a set of points.
func ConvexHull(points []Point) []Point {
	n := len(points)
	if n <= 1 {
		return points
	}

	// Sort points lexicographically (by x, then by y)
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	// Approach: construct lower and upper hulls
	lower := make([]Point, 0, n)
	for _, p := range points {
		// As long as there are at least two points, check orientation of the triplet
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1] // Remove the point that doesn't form a left turn
		}
		lower = append(lower, p)
	}

	upper := make([]Point, 0, n)
	for i := n - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1] // Remove the point that doesn't form a left turn
		}
		upper = append(upper, p)
	}

	// Concatenate lower and upper hull.
	// Last point of each list is omitted to avoid repetition of first point from the other list.
	hull := append(lower[:len(lower)-1], upper[:len(upper)-1]...)

	return hull
}

func main() {
	points := []Point{
		{0, 3}, {2, 2}, {1, 1}, {2, 1}, {3, 0},
		{0, 0}, {3, 3},
	}

	hull := ConvexHull(points)
	fmt.Println("Convex hull:")
	for _, p := range hull {
		fmt.Printf("(%d, %d)\n", p.X, p.Y)
	}
}
