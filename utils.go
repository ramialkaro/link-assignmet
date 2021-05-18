package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Generate a list of link station which is consist of the following X, Y, and R values
func GetLinkStation() []StationStruct {
	linkStations := []StationStruct{
		{x: 0, y: 0, reach: 3},
		{x: 20, y: 20, reach: 5},
		{x: 10, y: 0, reach: 12},
		{x: 17, y: 17, reach: 2},
		{x: 30, y: 34, reach: 6},
		{x: 41, y: 43, reach: 4},
		{x: 1, y: 3, reach: 4},
	}
	return linkStations
}

// Print all the link station information.
func PrintLinkStationInfo() {
	color.Cyan("List of link station:")
	for _, station := range GetLinkStation() {
		fmt.Printf("X: %v,\t\tY: %v,\t\tReach: %v\n", station.x, station.y, station.reach)
	}
}
