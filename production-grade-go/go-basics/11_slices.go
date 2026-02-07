package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println(len(s))
	fmt.Println(cap(s)) // capacity is the number of elements in the underlying array (starting from the first element in ths slice to the last element in the array)

	fmt.Println("----------------")

	x := []int{2, 3, 5, 7, 11, 13}

	x = x[1:4]
	fmt.Println(x)

	x = x[:2]
	fmt.Println(x)

	x = x[1:]
	fmt.Println(x)

	fmt.Println("----------------")

	y := []int{2, 3, 5, 7, 11, 13}
	printSlice(y)

	y = y[:0]
	printSlice(y)

	y = y[:4]
	printSlice(y)

	y = y[2:]
	printSlice(y)

	fmt.Println("----------------")

	var slice []int
	fmt.Println(slice, len(slice), cap(slice))
	if slice == nil {
		fmt.Println("zero value of a slice is nil")
	}

	fmt.Println("----------------")

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}