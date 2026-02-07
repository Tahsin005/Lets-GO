package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	func() {
		fmt.Println("This is an anonymous function executed immediately.")
	}()

	hypot := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}
	fmt.Printf("Result of anonymous hypotenuse: %f\n", hypot(5, 3))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}