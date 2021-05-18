package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Generate a list of link station which is consist of the following X, Y, and R values
func GetLinkStation() []StationStruct {
	linkStations := []StationStruct{
		{x: 0, y: 0, r: 3},
		{x: 20, y: 20, r: 5},
		{x: 10, y: 0, r: 12},
		{x: 17, y: 17, r: 2},
		{x: 30, y: 34, r: 6},
		{x: 41, y: 43, r: 4},
		{x: 1, y: 3, r: 4},
	}
	return linkStations
}

// Print all the link station information.
func PrintLinkStationInfo() {
	color.Cyan("List of link station:")
	for _, station := range GetLinkStation() {
		fmt.Printf("x: %v,\t\ty: %v,\t\tr: %v\n", station.x, station.y, station.r)
	}
}
