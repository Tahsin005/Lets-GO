package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s Shape) {
	fmt.Println("Area is:", s.Area())
}

// error interface
type CalculationError struct {
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{
			msg: "Invalid input, value cannot be negative",
		}
	}

	return  math.Sqrt(val), nil
}

func main() {
	rect := Rectangle{
		width: 5,
		height: 4,
	}
	circle := Circle{
		radius: 5,
	}

	calculateArea(rect)
	calculateArea(circle)

	// interface type
	mysteryBox1 := interface{}(10)
	describeValue(mysteryBox1)
	mysteryBox2 := interface{}(10.00)
	describeValue(mysteryBox2)
	mysteryBox3 := interface{}("10")
	describeValue(mysteryBox3)

	retrievedInt, ok := mysteryBox1.(int)
	if ok {
		fmt.Println("Retrieved value is an integer:", retrievedInt * 100)
	} else {
		fmt.Println("Value is not an integer")
	}

	performCalculation(81)
	performCalculation(-10)
}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}