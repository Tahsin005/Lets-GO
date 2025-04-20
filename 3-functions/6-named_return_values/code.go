package main

import "fmt"

func getCoords() (x, y int) {
	// x and y are initialized with zero

	// can update the variables if we want
	x += 8
	y += 9
	return // automaically return x and y
}

func yearsUntilEvent(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilRental int) {
	yearsUntilAdult = 18 - age
	yearsUntilDrinking = 21 - age
	yearsUntilRental = 25 - age

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	if yearsUntilRental < 0 {
		yearsUntilRental = 0
	}
	// return // this returns the values implicitly
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilRental // can return the variables by name or by value
}

func main () {
	first, second := getCoords()
	fmt.Println(first, second) 

	yua, yud, yur := yearsUntilEvent(23)
	fmt.Println("Years until adult:", yua)
	fmt.Println("Years until drinking:", yud)
	fmt.Println("Years until rental:", yur)
}