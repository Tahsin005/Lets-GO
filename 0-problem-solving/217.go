package main

import "fmt"

// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/description/?envType=daily-question&envId=2025-08-05

func maximumLength(nums []int, k int) int {
	ans := 2
	for i := 0; i < k; i++ {
		remainder := make([]int, k)
		for _, num := range nums {
			numMod := num % k
			requiredMod := (i - numMod + k) % k
			remainder[numMod] = remainder[requiredMod] + 1
		}

		for _, length := range remainder {
			if length > ans {
				ans = length
			}
		}
	}

	return ans
}

func main () {

}