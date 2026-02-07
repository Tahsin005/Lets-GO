package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// similar to while loop
	x := 1
	for x < 1000 {
		x += x
	}
	
	fmt.Println(x)

	// infinite loop
	// for {
	// 	fmt.Println("faaa")
	// }
}