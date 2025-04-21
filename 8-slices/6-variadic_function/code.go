package main

import "fmt"

// variadic function
func sum (nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main () {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	// spread operator
	nums := []int{1, 2, 3, 4, 5}
	total = sum(nums...)
	fmt.Println("Total using spread operator:", total)
}