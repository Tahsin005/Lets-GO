package main

import "fmt"

type comparable interface {
	int | float64 | string
}

func FindMax[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
    // Test with integers
    intMax := FindMax(3, 5)
    fmt.Println("Max of 3 and 5:", intMax) // Expected output: 5

    // Test with floats
    floatMax := FindMax(4.5, 3.7)
    fmt.Println("Max of 4.5 and 3.7:", floatMax) // Expected output: 4.5

    // Test with strings
    stringMax := FindMax("apple", "banana")
    fmt.Println("Max of 'apple' and 'banana':", stringMax) // Expected output: "banana"
}