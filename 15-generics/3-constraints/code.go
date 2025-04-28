package main

import "fmt"

// defining a custom constraints
type NumberOrString interface {
	int | float64 | string
}

func Max[T NumberOrString] (a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(10, 20))        // Works with int
	fmt.Println(Max(3.14, 2.71))     // Works with float64
	fmt.Println(Max("apple", "banana")) // Works with string
}