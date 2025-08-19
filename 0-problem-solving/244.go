package main

import "fmt"

// https://leetcode.com/problems/number-of-zero-filled-subarrays/description/?envType=daily-question&envId=2025-08-19

func zeroFilledSubarray(nums []int) int64 {
	var count, subcount int64 = 0, 0

	for _, num := range nums {
		if num != 0 {
			subcount = 0
			continue
		}

		subcount += 1
		count += subcount
	}

	return count
}

func main () {

}