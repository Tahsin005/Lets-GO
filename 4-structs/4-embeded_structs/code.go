package main

import "fmt"

type Car struct {
	model string
	madeBy string
}

type Truck struct {
	Car
	seatCapacity int
}

func main () {
	myTruch := Truck{
		seatCapacity: 2,
		Car: Car{
			model: "Cyber Truck v1",
			madeBy: "Tesla",
		},
	}

	// accesing the fields of the embedded struct
	fmt.Println("Truck Model: ", myTruch.model)
	fmt.Println("Truck Made By: ", myTruch.madeBy)
	fmt.Println("Truck Seat Capacity: ", myTruch.seatCapacity)
	fmt.Println("====================================")

	// accesing the fields of the embedded struct using the struct name
	// this is not necessary but can be used for clarity
	fmt.Println("Truck Model: ", myTruch.Car.model)
	fmt.Println("Truck Made By: ", myTruch.Car.madeBy)
	fmt.Println("Truck Seat Capacity: ", myTruch.seatCapacity)
}