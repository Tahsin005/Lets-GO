package main

import "fmt"

// https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/description/?envType=daily-question&envId=2025-06-28

func findKDistantIndices(nums []int, key int, k int) []int {
	result := make([]int, 0)
	right := 0

	for i, num := range nums {
		if num != key {
			continue
		}

		left := max(i - k, right)
		right = min(i + k + 1, len(nums))

		for j := left; j < right; j++ {
			result = append(result, j)
		}
	}
	return result
}


func main () {

}