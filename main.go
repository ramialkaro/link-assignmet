package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func main() {

	PrintLinkStationsInfo()

	var target TargetStruct
	var device DeviceStruct

	fmt.Println("\n\nhint, values should not be in string format, only numbers are accepted")

	fmt.Print("Enter X value for device: ")
	_, err := fmt.Scanf("%d", &device.x)

	// Throw error when the X value of the device is not number else keep going
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Print("Enter Y value for device: ")
		_, err = fmt.Scanf("%d", &device.y)

		// Throw error when the Y value of the device is not number else keep going
		if err != nil {
			log.Fatal(err)
		} else {

			// Go through all link-stations.
			for _, station := range GetLinkStations() {

				// calculate distance between device and station.
				distance := CalculateDistance(device, station)

				// make the power zero when distance bigger than reach.
				if distance > station.reach {
					if target.power == 0 {
						target.power = 0
					}
				} else {

					// read new power value
					nextPower := CalculatePower(station.reach, distance)

					// compare current power with new one.
					// getting the max power for certain link station and make it (most power)
					if target.power <= nextPower {
						target.power = nextPower
						target.x = station.x
						target.y = station.y
					}

				}
			}

			// show a message contain power, most power link-stationin for device point in console.
			if target.power != 0 {
				color.Green("\nBest link station for point x:%d, y:%d is x:%d, y:%d with power %d\n",
					device.x,
					device.y,
					target.x,
					target.y,
					target.power)

				// Not found a link station
			} else if target.power == 0 {
				fmt.Printf("\nNo link station within reach for point x: %v, y: %v \n", device.x, device.y)
			}
		}
	}

}
