package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a, a[0], a[1])

	arr := [2]string{"hello", "world"}
	fmt.Println(arr)

	arr2 := [2]int{1, 2}
	fmt.Println(arr2)
}