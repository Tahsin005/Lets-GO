package main

import "fmt"

func sum(startingPoint int, s ...int) int {
	sum := 0
	for i := startingPoint; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func main() {
	fmt.Println("Hello World", "Variadic Function")
	fmt.Println(sum(2, 1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(4, nums...))
}