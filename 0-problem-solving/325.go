package main

import "fmt"

// https://leetcode.com/problems/compute-decimal-representation/

func decimalRepresentation(n int) []int {
	var res []int
	i := 1

	for n > 0 {
		x := n % 10
		if x > 0 {
			x = x * i
			res = append([]int{x}, res...)
		}
		i = i * 10
		n = n / 10
	}

	return res
}

func main () {

}