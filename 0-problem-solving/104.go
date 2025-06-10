package main

import "fmt"

// https://leetcode.com/problems/split-the-array/description/

func isPossibleToSplit(nums []int) bool {
    freq := make([]int, 101)

    for _, n := range nums {
        freq[n]++
        if freq[n] >= 3 {
            return false
        }
    }

    return true
}

func main () {

}