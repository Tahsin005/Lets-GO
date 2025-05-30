package main

import "fmt"

// https://leetcode.com/problems/squares-of-a-sorted-array/description/?envType=problem-list-v2&envId=array

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n - 1

	for i := n - 1; i >= 0; i-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}

	return result
}


func main () {

}