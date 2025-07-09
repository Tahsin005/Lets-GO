package main

import "fmt"

// https://leetcode.com/problems/running-sum-of-1d-array/description/?envType=problem-list-v2&envId=array

func runningSum(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        nums[i] = nums[i] + nums[i - 1]
    }
    return nums
}

func main () {

}