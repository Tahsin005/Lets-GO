package main

import "fmt"

// https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/description/?envType=daily-question&envId=2025-07-25

func maxSum(nums []int) int {
    m := make(map[int]bool)
    for _,v := range nums {
        m[v] = true
    }
    var result int
    for i := 100; i > 0; i-- {
        if m[i] == true {
            result += i
        }
    }
    if result == 0 {
        for i := 0; i >= -100; i-- {
            if m[i] == true {
                result = i
                break
            }
        }
    }
    return result
}

func main () {

}