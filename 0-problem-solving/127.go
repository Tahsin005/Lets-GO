package main

import "fmt"

// https://leetcode.com/problems/sum-of-variable-length-subarrays/description/?envType=problem-list-v2&envId=array

func subarraySum(nums []int) int {
	prefixSum := make([]int, len(nums))

	var total int
	for i := 0; i < len(nums); i++ {
		sum := 0
		start := max(0, i - nums[i])
		for j := start; j <= i; j++ {
			sum += nums[j]
		}
		prefixSum[i] = sum
		total += prefixSum[i]
	}

	return total
}

func main () {

}