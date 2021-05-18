package main

import (
	"fmt"
)

func main() {
	var xDevice, yDevice int

	fmt.Println("hint, values should not be in string format, only numbers are accepted")
	fmt.Print("Enter X: ")
	_, _ = fmt.Scanf("%d", &xDevice)
	fmt.Print("Enter Y: ")
	_, _ = fmt.Scanf("%d", &yDevice)

	fmt.Printf("x: %v and y: %v\n", xDevice, yDevice)
	stations := GetLinkStation()
	GetResult(stations, xDevice, yDevice)

}
