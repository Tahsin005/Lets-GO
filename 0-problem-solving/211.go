package main

import "fmt"

// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/description/?envType=daily-question&envId=2025-08-02

func countHillValley(nums []int) int {
	j := 0
	res := 0
	for i := 1; i < len(nums) - 1; i++ {
        if nums[i] != nums[i - 1] {
            j = i - 1
        }
		if nums[i] > nums[j] && nums[i] > nums[i + 1] ||
			nums[i] < nums[j] && nums[i] < nums[i + 1] {
			res++
		}
	}

	return res
}

func main () {

}