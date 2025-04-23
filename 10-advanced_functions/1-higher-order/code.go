package main

import "fmt"

func add (a, b int) int {
	return a + b
}

func mul (a, b int)	int {
	return a * b
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main () {
	fmt.Println(aggregate(9, 2, 3, add)) // 6
	fmt.Println(aggregate(9, 2, 3, mul)) // 54
}