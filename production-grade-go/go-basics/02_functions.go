package main

import (
	"fmt"
	"math/rand"
)

func split(sum int) (x, y int) {
	x += 2 + sum
	y += 9 - sum
	return
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2)) 
	a, b := split(44)
	fmt.Println(a, b)
	fmt.Println("My favourite number is", rand.Intn(10))
}
