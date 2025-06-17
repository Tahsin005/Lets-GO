package main

import "fmt"

// https://leetcode.com/problems/maximum-difference-between-increasing-elements/description/?envType=daily-question&envId=2025-06-17

func maximumDifference(nums []int) int {
    min_val := nums[0]
    max_diff := -1
    for _, num := range nums[1:] {
        if num > min_val {
            diff := num - min_val
            if diff > max_diff {
                max_diff = diff
            }
        } else {
            min_val = num
        }
    }
    return max_diff
}

func main() {

}