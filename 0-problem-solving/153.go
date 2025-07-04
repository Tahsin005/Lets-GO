package main

import "fmt"

// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/description/?envType=problem-list-v2&envId=array

func countPairs(nums []int, target int) (res int) {
	sort.Ints(nums)
	for left, right := 0, len(nums) - 1; left < right; {
		if nums[left] + nums[right] < target {
            res += right - left
            left++
		} else {
		    right--
        }
	}
	return
}


func main () {

}