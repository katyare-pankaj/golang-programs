package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

const (
	basketHeight = 7.5 // Height of the basketball hoop in feet
)

func main() {
	// Free-throw line coordinates (x1, y1), (x2, y2)
	freeThrowLine := mat.NewDense(2, 2, []float64{22, 0, 22, 14})

	// Shot location (x, y)
	shotLocation := mat.NewDense(1, 2, []float64{22, 10})

	// Calculate the optimal shooting angle
	theta := optimalShootingAngle(freeThrowLine, shotLocation)
	fmt.Println("Optimal shooting angle:", theta*180/math.Pi, "degrees")
}

func optimalShootingAngle(freeThrowLine, shotLocation *mat.Dense) float64 {
	// Calculate the unit vector of the free-throw line
	dir := mat.NewDense(2, 1, nil)
	freeThrowLine.ColView(0, dir)
	dir.Sub(freeThrowLine.ColView(1, nil), dir)
	dir.DivScalar(dir.Norm(mat.Norm2), dir)

	// Calculate the angle between the shot direction and the free-throw line
	dot := mat.Dot(shotLocation, dir)
	shotMagnitude := shotLocation.Norm(mat.Norm2)
	lineMagnitude := dir.Norm(mat.Norm2)
	cosTheta := dot / (shotMagnitude * lineMagnitude)
	theta := math.Acos(cosTheta)

	// Adjust the angle for the quadrant
	if shotLocation.At(0, 1) < 0 {
		theta = math.Pi - theta
	}

	return theta
}
