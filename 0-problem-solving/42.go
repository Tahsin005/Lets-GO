package main

import "fmt"

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/?envType=problem-list-v2&envId=array

func smallerNumbersThanCurrent(nums []int) []int {
    ln := len(nums)
    b := make([]int, ln)
    for i := 0; i < ln; i++ {
        c := 0
        for j := 0; j < ln; j++ {
            if nums[i] > nums[j] { c++ }
        }
        b[i] = c
    }
    return b
}

func main () {

}