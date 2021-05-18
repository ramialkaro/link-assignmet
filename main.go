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

	fmt.Print("Enter X value for device: ")
	_, err := fmt.Scanf("%d", &device.x)

	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Print("Enter Y value for device: ")
		_, err = fmt.Scanf("%d", &device.y)

		if err != nil {
			log.Fatal(err)
		} else {

			for _, station := range GetLinkStation() {

				distance := CalculateDistance(device, station)

				if distance > station.reach {
					if target.power == 0 {
						target.power = 0
					}
				} else {
					nextPower := CalculatePower(station.reach, distance)
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
				fmt.Printf("\nNo link station within reach for point x: %v, y: %v \n", device.x, device.y)
			}
		}
	}

}
