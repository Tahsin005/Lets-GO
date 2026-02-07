package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	defer fmt.Println("hello - 2")
	fmt.Println("hello - 1")

	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
}