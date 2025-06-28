package main

import "fmt"

type Number interface {
	~int | ~float64 // Matches int, float64, and any named types with underlying int or float64
}

type MyInt int // MyInt has an underlying type of int

func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	var x MyInt = 10
	var y MyInt = 20
	sum := Add(x, y) // This works because MyInt's underlying type is int
	fmt.Println(sum)

	var a int = 10
	var b int = 20
	sum2 := Add(a, b)
	fmt.Println(sum2)
}