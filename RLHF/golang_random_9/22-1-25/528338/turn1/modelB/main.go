package main

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y float64
}
type byPolarAngle []Point

func (a byPolarAngle) Len() int      { return len(a) }
func (a byPolarAngle) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byPolarAngle) Less(i, j int) bool {
	dx1, dy1 := a[i].X-a[0].X, a[i].Y-a[0].Y
	dx2, dy2 := a[j].X-a[0].X, a[j].Y-a[0].Y
	// Check if they have the same polar angle
	if dy1*dx2 == dy2*dx1 {
		// If yes, then compare the distances
		return dx1*dx1+dy1*dy1 < dx2*dx2+dy2*dy2
	}
	// Otherwise, compare the polar angles
	return dy1*dx2 < dy2*dx1
}
func ConvexHull(points []Point) []Point {
	// Find the point with the minimum y-coordinate
	minIdx := 0
	for i := 1; i < len(points); i++ {
		if points[i].Y < points[minIdx].Y {
			minIdx = i
		}
	}
	// Swap the minimum point to the first index
	points[0], points[minIdx] = points[minIdx], points[0]
	// Sort the remaining points by polar angle around the first point
	sort.Sort(byPolarAngle(points[1:]))
	stack := []Point{points[0], points[1]}
	for i := 2; i < len(points); i++ {
		p := points[i]
		q := stack[len(stack)-1]
		r := stack[len(stack)-2]
		// Check if the current point p is on the left side of the line qr
		if (q.X-r.X)*(p.Y-r.Y) > (q.Y-r.Y)*(p.X-r.X) {
			stack = append(stack, p)
		}
	}
	return stack
}
func main() {
	points := []Point{
		Point{X: 1, Y: 1},
		Point{X: 3, Y: 3},
		Point{X: 2, Y: 2},
		Point{X: 5, Y: 5},
		Point{X: 4, Y: 1},
		Point{X: 0, Y: 4},
	}
	hull := ConvexHull(points)
	fmt.Println("Convex Hull:")
	for _, p := range hull {
		fmt.Printf("(%v, %v)\n", p.X, p.Y)
	}
}
