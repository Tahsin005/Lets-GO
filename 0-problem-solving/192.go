package main

import "fmt"

// https://leetcode.com/problems/maximum-erasure-value/description/?envType=daily-question&envId=2025-07-24

func maximumUniqueSubarray(nums []int) int {
    m := make(map[int]int)
    start := 0
    sum := 0
    maxSum := 0
    for i := 0; i < len(nums); i++ {
        if v, ok := m[nums[i]]; ok {
            for j := start; j < v; j++ {
                sum -= nums[j]
                delete(m, nums[j])
            }
            start = v+1
        } else {
            sum += nums[i]
        }
        m[nums[i]] = i
        maxSum = max(maxSum, sum)
    }
    return maxSum
}

func main () {

}