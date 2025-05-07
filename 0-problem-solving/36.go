package main

import "fmt"

// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/description/?envType=problem-list-v2&envId=array

func sumOddLengthSubarrays(arr []int) int {
    res := 0
    n := len(arr)
    for i, num := range arr {
        cnt := ((i + 1) * (n - i) + 1) / 2
        res += num * cnt
    }
    return res
}

func main () {

}