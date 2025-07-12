package main

import "fmt"

// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/description/?envType=problem-list-v2&envId=array

func minOperations(nums []int, k int) int {
    var sum int
    for _,v := range nums {
        sum += v
    }
    return sum % k
}

func main () {

}