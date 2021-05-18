package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func main() {

	PrintLinkStationInfo()

	var target TargetStruct
	var device DeviceStruct

	fmt.Println("\n\nhint, values should not be in string format, only numbers are accepted")

	// read X value of the device from console.
	fmt.Print("Enter X: ")
	_, err := fmt.Scanf("%d", &device.x)

	if err != nil {
		log.Fatal(err)
	} else {

		// read Y value of the device from console
		fmt.Print("Enter Y: ")
		_, err = fmt.Scanf("%d", &device.y)

		if err != nil {
			log.Fatal(err)
		} else {

			for _, station := range GetLinkStation() {

				distance := CalculateDistance(device, station)

				if distance > station.r {
					if target.power == 0 {
						target.power = 0
					}
				} else {
					nextPower := CalculatePower(station.r, distance)
					if target.power <= nextPower {
						target.power = nextPower
						target.x = station.x
						target.y = station.y
					}
				}
			}

			if target.power != 0 {
				color.Green("\nBest link station for point x:%d, y:%d is x:%d, y:%d with power %d\n",
					device.x,
					device.y,
					target.x,
					target.y,
					target.power)

			} else if target.power == 0 {
				fmt.Printf("\nNo link station with reach for point x: %v, y: %v \n", device.x, device.y)
			}
		}
	}

}
