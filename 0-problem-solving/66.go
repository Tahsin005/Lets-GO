package main

import "fmt"

// https://leetcode.com/problems/type-of-triangle/description/?envType=daily-question&envId=2025-05-22

func triangleType(nums []int) string {
	a, b, c := nums[0], nums[1], nums[2]
	if a + b <= c || a + c <= b || b + c <= a {
		return "none"
	}
	if a == b && b == c {
		return "equilateral"
	}
	if a == b {
		return "isosceles"
	}
	if b == c {
		return "isosceles"
	}
	if a == c {
		return "isosceles"
	}
	return "scalene"
}

func main () {

}