package main

import "fmt"

// https://leetcode.com/problems/minimum-operations-to-equalize-array/

func minOperations(nums []int) int {
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] != nums[i - 1] {
            return 1
        }
    }
    return 0
}

func main () {

}