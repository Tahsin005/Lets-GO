package main

import "fmt"

// https://leetcode.com/problems/create-target-array-in-the-given-order/description/?envType=problem-list-v2&envId=array

func createTargetArray(nums []int, index []int) []int {
    target := make([]int, len(nums))

    for idx, i := range index {
        for j := len(target) - 1; j > i; j-- {
            target[j] = target[j - 1]
        }

        target[i] = nums[idx]
    }

    return target
}

func main() {
	
}