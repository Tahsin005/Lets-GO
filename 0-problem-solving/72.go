package main

import "fmt"

// https://leetcode.com/problems/separate-the-digits-in-an-array/description/?envType=problem-list-v2&envId=array

func separateDigits(nums []int) []int {
	var result []int
	var digits [6]int
	for _, num := range nums {
		i := 5
		for ; num != 0; i, num = i - 1, num / 10 {
			digits[i] = num % 10
		}
		result = append(result, digits[i + 1:]...)
	}
	return result
}

func main () {

}