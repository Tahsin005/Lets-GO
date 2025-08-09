package main

import "fmt"

// https://leetcode.com/problems/trionic-array-i/description/

func isTrionic(nums []int) bool {
	if nums[1] <= nums[0] {
		return false
	}
	state := 1
	for i := 2; i < len(nums); i++ {
		if nums[i - 1] == nums[i] {
			return false
		}

		if state == 1 {
			if nums[i - 1] < nums[i] {
				continue
			} else {
				state = 2
			}
		} else if state == 2 {
			if nums[i - 1] > nums[i] {
				continue
			} else {
				state = 3
			}
		} else {
			if nums[i - 1] < nums[i] {
				continue
			} else {
				return false
			}
		}
	}
	return state == 3
}

func main () {

}