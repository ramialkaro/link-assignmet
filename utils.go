package main

import (
	"fmt"

	"github.com/fatih/color"
)

// List of link station which is consist of the following X, Y, R
func GetLinkStation() [][3]int {
	stations := [][3]int{
		{0, 0, 3},
		{20, 20, 5},
		{10, 0, 12},
		{17, 17, 2},
		{30, 34, 6},
		{41, 43, 4},
		{1, 3, 4},
	}
	return stations
}

// validate the distance between device location from the list of the link station
// args of the function are list of links station([x, y, r]) and device point(x, y)
// s[0] is x value for station,
// s[1] is y value for station,
// s[2] is reach value for station

func GetResult(stations [][3]int, xDevice int, yDevice int) {

	var power int
	var x int
	var y int

	color.Cyan("\nList of link stations:")
	for _, s := range stations {

		distance := CalculateDistance(xDevice, yDevice, s[0], s[1])

		fmt.Printf("x: %v,\ty: %v,\tr%v\t\tdistance: %v\n", s[0], s[1], s[2], distance)
		//fmt.Printf("loop x: %d, y: %d, r:%d\tdistance: %d\n\n", s[0], s[1], s[2], distance)

		if distance > s[2] {
			if power == 0 {
				power = 0
			}
		} else {
			nextPower := CalculatePower(s[2], distance)
			if power <= nextPower {
				power = nextPower
				x = s[0]
				y = s[1]
			}
			//fmt.Printf("Power value is %d\n", power)
		}

	}

	if power != 0 {
		color.Green("\nBest link station for point x:%d, y:%d is x:%d, y:%d with power %d\n",
			xDevice,
			yDevice,
			x,
			y,
			power)

	} else if power == 0 {
		fmt.Printf("\nNo link station with reach for point x: %v, y: %v \n", xDevice, yDevice)
	}
}
