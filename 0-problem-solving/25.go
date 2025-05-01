package main

import "fmt"

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/?envType=daily-question&envId=2025-04-25

func findNumbers(nums []int) int {
	result := 0
	for _, num := range nums {
		if countDigits(num) % 2 == 0 {
			result++
		}
	}
	return result
}

func countDigits(num int) int {
	if num <= 0 {
		panic("problem constraints have been changed")
	}

	result := 0
	for num > 0 {
		num = num / 10
		result++
	}
	return result
}


func main () {

}