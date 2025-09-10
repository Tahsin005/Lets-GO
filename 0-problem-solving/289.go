package main

import "fmt"

// https://leetcode.com/problems/sum-of-even-numbers-after-queries/description/

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    evenSum := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 == 0 {
            evenSum += nums[i];
        }
    }

    var ans []int
    for i := 0; i < len(queries); i++ {
        val, index := queries[i][0], queries[i][1]

        if nums[index] % 2 == 0 {
            evenSum -= nums[index]
        }

        nums[index] += val
        if nums[index] % 2 == 0 {
            evenSum += nums[index]
        }
        ans = append(ans, evenSum)
    }

    return ans
}

func main () {

}