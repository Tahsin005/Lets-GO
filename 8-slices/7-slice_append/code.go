package main

import "fmt"

func main () {
	slice := []int{}

	// append 1 item -- append is a variadic function
	slice = append(slice, 1)
	fmt.Println(slice) // [1]

	// append 2 items
	slice = append(slice, 2, 3)
	fmt.Println(slice) // [1 2 3]

	secondSlice := []int{4, 5}
	// append 2 slices
	slice = append(slice, secondSlice...)
	fmt.Println(slice) // [1 2 3 4 5]
}