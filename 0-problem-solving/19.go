package main

import "fmt"

// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/description/?envType=daily-question&envId=2025-04-27

func countSubarrays(nums []int) int {
    count := 0
    if len(nums) < 3 {
        return 0
    }

    left, mid, right := 0, 1, 2

    for right < len(nums) {
        if 2 * (nums[left] + nums[right]) == nums[mid] {
            count++
        }
        left++
        mid++
        right++
    }

    return count
}

func main () {
	
}