package main

import "fmt"

// https://leetcode.com/problems/transform-array-by-parity/description/?envType=problem-list-v2&envId=array

func transformArray(nums []int) []int {
    result := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] % 2 == 0 {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}
	}

	sort.Ints(result)
	return result
}

func main() {

}