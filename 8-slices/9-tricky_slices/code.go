package main

import "fmt"

func main () {
	a := make([]int, 3)

	fmt.Println(a)

	a = append(a, 4)

	b := append(a, 4) // is this saying take whats in a and append 4 to it and then assign it to b?
	fmt.Println(b)

	c := append(a, 5)
	fmt.Println(c)

	fmt.Println(a)
}