package main

import (
	"fmt"
	"math"
)

// GeoPoint represents a geospatial point with longitude and latitude.
type GeoPoint struct {
	Longitude float64
	Latitude  float64
}

// QuadTreeNode represents a node in the quadtree.
type QuadTreeNode struct {
	MinLongitude float64
	MaxLongitude float64
	MinLatitude  float64
	MaxLatitude  float64
	Points       []GeoPoint
	Children     []*QuadTreeNode
}

// QuadTree represents the entire quadtree structure.
type QuadTree struct {
	Root *QuadTreeNode
}

// NewQuadTree creates a new quadtree.
func NewQuadTree(minLongitude, maxLongitude, minLatitude, maxLatitude float64) *QuadTree {
	return &QuadTree{
		Root: &QuadTreeNode{
			MinLongitude: minLongitude,
			MaxLongitude: maxLongitude,
			MinLatitude:  minLatitude,
			MaxLatitude:  maxLatitude,
		},
	}
}

// Insert adds a geospatial point to the quadtree.
func (qt *QuadTree) Insert(point GeoPoint) {
	qt.insertPoint(qt.Root, point)
}

func (qt *QuadTree) insertPoint(node *QuadTreeNode, point GeoPoint) {
	// If point is outside bounds of the current node, return
	if point.Longitude < node.MinLongitude || point.Longitude > node.MaxLongitude || point.Latitude < node.MinLatitude || point.Latitude > node.MaxLatitude {
		return
	}
	// If node contains no points, add the point and return
	if node.Points == nil {
		node.Points = []GeoPoint{point}
		return
	}
	// If node already contains points, see if we need to create children
	node.Points = append(node.Points, point)
	if len(node.Points) > 4 {
		qt.splitNode(node)
	}
}

func (qt *QuadTree) splitNode(node *QuadTreeNode) {
	middleLong := (node.MinLongitude + node.MaxLongitude) / 2
	middleLat := (node.MinLatitude + node.MaxLatitude) / 2

	ne := &QuadTreeNode{
		MinLongitude: node.MinLongitude,
		MaxLongitude: middleLong,
		MinLatitude:  middleLat,
		MaxLatitude:  node.MaxLatitude,
	}
	nw := &QuadTreeNode{
		MinLongitude: middleLong,
		MaxLongitude: node.MaxLongitude,
		MinLatitude:  middleLat,
		MaxLatitude:  node.MaxLatitude,
	}
	se := &QuadTreeNode{
		MinLongitude: node.MinLongitude,
		MaxLongitude: middleLong,
		MinLatitude:  node.MinLatitude,
		MaxLatitude:  middleLat,
	}
	sw := &QuadTreeNode{
		MinLongitude: middleLong,
		MaxLongitude: node.MaxLongitude,
		MinLatitude:  node.MinLatitude,
		MaxLatitude:  middleLat,
	}

	node.Children = []*QuadTreeNode{ne, nw, se, sw}

	for _, point := range node.Points {
		qt.insertPoint(ne, point)
		qt.insertPoint(nw, point)
		qt.insertPoint(se, point)
		qt.insertPoint(sw, point)
	}

	node.Points = nil
}

// SearchByProximity searches the quadtree for points within a given radius of a reference point.
func (qt *QuadTree) SearchByProximity(referencePoint GeoPoint, radius float64) ([]GeoPoint, error) {
	// Extract the possible bounding box for the points within the specified radius.
	// This bounding box will be used to decide which nodes in the quadtree to search.
	minLongitude := referencePoint.Longitude - radius
	maxLongitude := referencePoint.Longitude + radius
	minLatitude := referencePoint.Latitude - radius
	maxLatitude := referencePoint.Latitude + radius

	// Start the search at the root node
	return qt.searchNodesByProximity(qt.Root, referencePoint, radius, minLongitude, maxLongitude, minLatitude, maxLatitude)
}

func (qt *QuadTree) searchNodesByProximity(node *QuadTreeNode, referencePoint GeoPoint, radius float64, minLongitude, maxLongitude, minLatitude, maxLatitude float64) ([]GeoPoint, error) {
	var points []GeoPoint

	// If point is within this node, add all points in this node
	if !(referencePoint.Longitude < node.MinLongitude || referencePoint.Longitude > node.MaxLongitude || referencePoint.Latitude < node.MinLatitude || referencePoint.Latitude > node.MaxLatitude) {
		if node.Points != nil {
			for _, point := range node.Points {
				if qt.pointsWithinRadius(referencePoint, point, radius) {
					points = append(points, point)
				}
			}
		}
	}

	// Search child nodes
	for _, child := range node.Children {
		childPoints, err := qt.searchNodesByProximity(child, referencePoint, radius, minLongitude, maxLongitude, minLatitude, maxLatitude)
		if err != nil {
			return nil, err
		}
		points = append(points, childPoints...)
	}

	return points, nil
}

// pointsWithinRadius checks if two points are within a specified radius of each other.
func (qt *QuadTree) pointsWithinRadius(point1, point2 GeoPoint, radius float64) bool {
	// A simple approximation using the haversine formula
	lat1 := math.Radians(point1.Latitude)
	lon1 := math.Radians(point1.Longitude)
	lat2 := math.Radians(point2.Latitude)
	lon2 := math.Radians(point2.Longitude)

	Δlon := lon2 - lon1
	Δlat := lat2 - lat1

	a := math.Sin(Δlat/2)*math.Sin(Δlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(Δlon/2)*math.Sin(Δlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return math.Radians(radius) >= c
}

func main() {
	// Example usage
	points := []GeoPoint{
		{Longitude: 4.8952, Latitude: 52.3702},
		{Longitude: 4.8909, Latitude: 52.3759},
		{Longitude: 5.2212, Latitude: 52.3736},
		{Longitude: -1.2852, Latitude: 51.4776},
		{Longitude: -0.1278, Latitude: 51.5074},
		{Longitude: -73.5015, Latitude: 40.7823},
		{Longitude: -74.0060, Latitude: 40.7128},
	}

	minLongitude := math.Min(points...)
	maxLongitude := math.Max(points...)
	minLatitude := math.Min(points...)
	maxLatitude := math.Max(points...)

	qt := NewQuadTree(minLongitude, maxLongitude, minLatitude, maxLatitude)
	for _, point := range points {
		qt.Insert(point)
	}

	radius := 5.0 // Radius in degrees
	referencePoint := GeoPoint{Longitude: 0.0, Latitude: 0.0}

	groupedPoints, err := qt.SearchByProximity(referencePoint, radius)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Points within", radius, "degrees of (0.0, 0.0):")
	for _, point := range groupedPoints {
		fmt.Printf("(%.4f, %.4f)\n", point.Longitude, point.Latitude)
	}
}
