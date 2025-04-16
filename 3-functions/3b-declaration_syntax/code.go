package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func f(cb func(a int, b int) int, c int) int {
	result := cb(10, c)
	return result
}

func main () {
	fmt.Println("Hurray!!")
	resultOfFunction := f(add, 8)
	fmt.Println(resultOfFunction)
}