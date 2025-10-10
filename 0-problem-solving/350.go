package main

import "fmt"

// https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/

func longestSubsequence(nums []int) int {
    xorSum := 0
    allZero := true
    for _, num := range nums {
        xorSum ^= num
        if num != 0 {
            allZero = false
        }
    }
    if allZero {
        return 0
    }
    if xorSum != 0 {
        return len(nums)
    }
    return len(nums) - 1
}

func main () {

}