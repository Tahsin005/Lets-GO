package main

import "fmt"

// https://leetcode.com/problems/set-mismatch/description/?envType=problem-list-v2&envId=array

func findErrorNums(nums []int) []int {
    n := len(nums)
    numMap := make(map[int]bool)
    duplicate := -1
    actualSum := 0

    for _, num := range nums {
        if numMap[num] {
            duplicate = num
        }
        numMap[num] = true
        actualSum += num
    }

    expectedSum := (n * (n + 1)) / 2
    missingNum := expectedSum - (actualSum - duplicate)

    return []int{duplicate, missingNum}
}

func main () {

}