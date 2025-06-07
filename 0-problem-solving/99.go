package main

import "fmt"

// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/?envType=problem-list-v2&envId=array

func maxProductDifference(nums []int) int {
	sort.Ints(nums)

	left := nums[0] * nums[1]
	rigth := nums[len(nums) - 1] * nums[len(nums) - 2]

	res := rigth - left

	return res
}

func main () {

}