package main

import (
	"fmt"
	"math"
)

func main() {

	// justRight := 88
	// var temperature int
	// fmt.Println("What temperature is the soup?")
	// result, err := fmt.Scan(&temperature)

	// if err != nil {
	// 	fmt.Println("Error reading temperature:", err)
	// }

	// fmt.Println("Result of scan: ", result)

	// if temperature == justRight {
	// 	fmt.Println("The soup is just right")
	// }

	fmt.Printf("%.48f\n", math.Pi)
	fmt.Printf("%v\n", math.Pi)
	fmt.Printf("%T", math.Pi)

	radius := 10
	fmt.Printf("\n Radius is a %T", radius)

	area := math.Pi * float64(radius*radius)

	fmt.Printf("\n Area is %v", area)
}
