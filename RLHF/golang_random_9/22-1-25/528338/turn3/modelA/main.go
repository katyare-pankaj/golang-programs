package main

import (
	"fmt"
	"math"
	"sort"
)

// Point represents a point in 2D space.
type Point struct {
	X, Y float64
}

// Unique filters out duplicate points from a slice of points.
func Unique(points []Point) []Point {
	uniquePoints := make(map[Point]struct{}, len(points))
	var result []Point
	for _, p := range points {
		if _, exists := uniquePoints[p]; !exists {
			uniquePoints[p] = struct{}{}
			result = append(result, p)
		}
	}
	return result
}

// Sort by polar angle with a reference to the base point.
type byPolarAngle struct {
	points []Point
	base   Point
}

func (a byPolarAngle) Len() int      { return len(a.points) }
func (a byPolarAngle) Swap(i, j int) { a.points[i], a.points[j] = a.points[j], a.points[i] }
func (a byPolarAngle) Less(i, j int) bool {
	// Calculate polar angles
	dx1, dy1 := a.points[i].X-a.base.X, a.points[i].Y-a.base.Y
	dx2, dy2 := a.points[j].X-a.base.X, a.points[j].Y-a.base.Y
	angle1 := math.Atan2(dy1, dx1)
	angle2 := math.Atan2(dy2, dx2)

	if angle1 == angle2 {
		// Same polar angle, use distance to break tie
		return dx1*dx1+dy1*dy1 < dx2*dx2+dy2*dy2
	}
	return angle1 < angle2
}

func cross(o, a, b Point) float64 {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

func ConvexHull(points []Point) []Point {
	if len(points) < 3 {
		return points
	}

	// Remove duplicates
	points = Unique(points)
	if len(points) < 3 {
		return points
	}

	// Find the point with the minimum y-coordinate (and min x if tie)
	minIdx := 0
	for i := 1; i < len(points); i++ {
		if points[i].Y < points[minIdx].Y || (points[i].Y == points[minIdx].Y && points[i].X < points[minIdx].X) {
			minIdx = i
		}
	}

	// Swap the minimum point to the first index
	points[0], points[minIdx] = points[minIdx], points[0]
	base := points[0]

	// Sort the points by polar angle with the base point
	sort.Sort(byPolarAngle{points[1:], base})

	// Building the convex hull
	stack := []Point{points[0], points[1]}
	for i := 2; i < len(points); i++ {
		for len(stack) >= 2 && cross(stack[len(stack)-2], stack[len(stack)-1], points[i]) <= 0 {
			stack = stack[:len(stack)-1] // Pop the top element
		}
		stack = append(stack, points[i])
	}

	return stack
}

func main() {
	points := []Point{
		{X: 0, Y: 3}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 4, Y: 4},
		{X: 0, Y: 0}, {X: 1, Y: 2}, {X: 3, Y: 1}, {X: 3, Y: 3},
		{X: 1, Y: 1}, {X: 1, Y: 1}, // Duplicate points
	}

	hull := ConvexHull(points)

	fmt.Println("Convex Hull:")
	for _, p := range hull {
		fmt.Printf("(%v, %v)\n", p.X, p.Y)
	}
}
