package main

import (
	"sort"
)

// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i/description/?envType=problem-list-v2&envId=array

func minOperations(nums []int, k int) int {
    sort.Ints(nums)
    ans := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < k {
            ans++
        } else {
            break
        }
    }

    return ans
}

func main () {

}