package main

import (
	"sort"
)

// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/description/?envType=daily-question&envId=2025-06-29

type Item struct {
	idx int
	val int
}

func maxSubsequence(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	arr := make([]Item, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[i] = Item{idx: i, val: nums[i]}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val > arr[j].val
	})

	arr = arr[:k]
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].idx < arr[j].idx
	})

	for i := 0; i < len(arr); i++ {
		nums[i] = arr[i].val
	}

	return nums[:k]
}

func main () {

}