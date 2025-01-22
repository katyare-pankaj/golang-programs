package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y float64
}
type byPolarAngle struct {
	points []Point
	base   Point
}

func (a byPolarAngle) Len() int      { return len(a.points) }
func (a byPolarAngle) Swap(i, j int) { a.points[i], a.points[j] = a.points[j], a.points[i] }
func (a byPolarAngle) Less(i, j int) bool {
	dx1, dy1 := a.points[i].X-a.base.X, a.points[i].Y-a.base.Y
	dx2, dy2 := a.points[j].X-a.base.X, a.points[j].Y-a.base.Y
	angle1 := math.Atan2(dy1, dx1)
	angle2 := math.Atan2(dy2, dx2)
	if angle1 == angle2 {
		return dx1*dx1+dy1*dy1 < dx2*dx2+dy2*dy2
	}
	return angle1 < angle2
}
func convexHull(points []Point) []Point {
	if len(points) < 3 {
		return points
	}
	minIdx := 0
	for i := 1; i < len(points); i++ {
		if points[i].Y < points[minIdx].Y || (points[i].Y == points[minIdx].Y && points[i].X < points[minIdx].X) {
			minIdx = i
		}
	}
	points[0], points[minIdx] = points[minIdx], points[0]
	base := points[0]
	sort.Sort(byPolarAngle{points[1:], base})
	k := 1
	for i := 1; i < len(points); i++ {
		for i < len(points)-1 && cross(base, points[i], points[i+1]) == 0 {
			i++
		}
		points[k] = points[i]
		k++
	}
	points = points[:k]
	stack := []Point{points[0], points[1]}
	for i := 2; i < len(points); i++ {
		for len(stack) >= 2 && cross(stack[len(stack)-2], stack[len(stack)-1], points[i]) <= 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, points[i])
	}
	return stack
}
func cross(o, a, b Point) float64 {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

func main() {
	points := []Point{
		Point{X: 1, Y: 1},
		Point{X: 3, Y: 3},
		Point{X: 2, Y: 2},
		Point{X: 5, Y: 5},
		Point{X: 4, Y: 1},
		Point{X: 0, Y: 4},
		// Duplicate points
		Point{X: 3, Y: 3},
		Point{X: 2, Y: 2},
	}
	hull := convexHull(points)
	fmt.Println("Convex Hull:")
	for _, p := range hull {
		fmt.Printf("(%v, %v)\n", p.X, p.Y)
	}
}
