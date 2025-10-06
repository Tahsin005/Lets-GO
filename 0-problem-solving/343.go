package main

import "fmt"

// https://leetcode.com/problems/compute-alternating-sum/description/

func alternatingSum(nums []int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        if i % 2 == 0 {
            count += nums[i]
        } else {
            count -= nums[i]
        }
    }
    return count
}

func main () {

}