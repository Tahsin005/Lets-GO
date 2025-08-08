package main

import "fmt"

// https://leetcode.com/problems/partition-array-into-two-equal-product-subsets/description/

func checkEqualPartitions(nums []int, target int64) bool {
    mul := int64(1)
    for _, n := range nums {
        mul *= int64(n)
    }

    if mul != target * target {
        return false
    }

    var bruteForce func(index int, curr int64) bool
    bruteForce = func(index int, curr int64) bool {
        if curr == target {
            return true
        }
        if curr > target {
            return false
        }
        if index == len(nums) {
            return false
        }

        take := bruteForce(index + 1, curr * int64(nums[index]))
        notTake := bruteForce(index + 1, curr)

        return take || notTake
    }

    return bruteForce(0, 1)
}

func main () {

}