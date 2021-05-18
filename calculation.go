package main

import "math"

// calculate how far a device from link station
func CalculateDistance(device DeviceStruct, station StationStruct) int {
	var distance DistanceStruct

	distance.x = (station.x - device.x)
	distance.y = (station.y - device.y)
	return int(
		math.Sqrt(
			float64(distance.x*distance.x) + float64(distance.y*distance.y)))
}

// Calculate the power, by give a reach of link station and distance as arguments.
func CalculatePower(reach int, distance int) int {
	return (reach - distance) * (reach - distance)

}
