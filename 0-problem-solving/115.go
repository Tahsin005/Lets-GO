package main

import "fmt"

// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/description/?envType=daily-question&envId=2025-06-15

func maxAdjacentDistance(nums []int) int {
    maxDiff := abs(nums[0] - nums[len(nums) - 1])
    for i := 1; i < len(nums); i++ {
        maxDiff = max(maxDiff, abs(nums[i] - nums[i - 1]))
    }
    return maxDiff
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func main() {

}