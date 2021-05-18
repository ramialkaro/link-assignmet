package main

import "math"

// calculate how far a device from link station
// Device points are x1, y1
// Link station points are	x2, y2
func CalculateDistance(x1 int, y1 int, x2 int, y2 int) int {
	xDist := (x2 - x1)
	yDist := (y2 - y1)
	return int(
		math.Sqrt(
			float64(xDist*xDist) + float64(yDist*yDist)))
}

// Calculate power function by give a reach and distance as arguments.
func CalculatePower(reach int, distance int) int {
	return (reach - distance) * (reach - distance)

}
