package main

import "fmt"

// https://leetcode.com/problems/permutations/description/?envType=problem-list-v2&envId=array

func permute(nums []int) [][]int {
    res := [][]int{}

    if len(nums) == 1 {
        return append(res, []int{nums[0]})
    }

    for i := 0; i < len(nums); i++ {
        nums[0], nums[i] = nums[i], nums[0]

        for _, tail := range(permute(nums[1:])) {
            res = append(res, append(tail, nums[0]))
        }

        nums[i], nums[0] = nums[0], nums[i]
    }

    return res
}

func main () {

}