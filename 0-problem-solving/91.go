package main

import "fmt"

// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/description/?envType=problem-list-v2&envId=array

func differenceOfSum(nums []int) int {
    res := 0
    for _, v := range nums {
        res += v
        for v > 0 {
            res -= v % 10
            v = v/10
        }
    }

    return res
}

func main () {

}