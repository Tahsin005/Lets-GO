package main

import "fmt"

// https://leetcode.com/problems/sum-of-squares-of-special-elements/description/?envType=problem-list-v2&envId=array

func sumOfSquares(nums []int) int {
    res := 0
    l := len(nums)
    for ind := 1; ind < l + 1; ind++ {
        if l % ind == 0 {
            res += nums[ind - 1] * nums[ind - 1]
        }
    }
    return res
}

func main () {

}