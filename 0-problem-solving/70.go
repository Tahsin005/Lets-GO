package main

import "fmt"

// https://leetcode.com/problems/number-of-arithmetic-triplets/description/?envType=problem-list-v2&envId=array

func arithmeticTriplets(nums []int, diff int) int {
    occ := make([]int, 201)

    for i := 0; i < len(nums); i++ {
        occ[nums[i]] = 1
    }
    pairs := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] + diff <= 200 && occ[nums[i] + diff] == 1 {
            if nums[i] + 2 * diff <= 200 && occ[nums[i] + 2 * diff] == 1 {
                pairs +=1
            }
        }
    }

    return pairs
}

func main () {

}