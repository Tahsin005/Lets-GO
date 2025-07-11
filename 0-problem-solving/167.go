package main

import "fmt"

// https://leetcode.com/problems/subsets/description/?envType=problem-list-v2&envId=array

func subsets(nums []int) [][]int {
	var result [][]int
	var curr []int

	var explore func(int)
	explore = func(index int) {
		if index == len(nums) {
			subset := make([]int, len(curr))
			copy(subset, curr)
			result = append(result, subset)
			return
		}

		curr = append(curr, nums[index])
		explore(index + 1)
		curr = curr[:len(curr) - 1]

		explore(index + 1)
	}

	explore(0)
	return result
}

func main () {

}