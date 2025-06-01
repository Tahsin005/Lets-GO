package main

import "fmt"

// https://leetcode.com/problems/keep-multiplying-found-values-by-two/description/?envType=problem-list-v2&envId=array

func findFinalValue(nums []int, original int) int {
    var numSet [1001]int

    for _, num := range nums {
        numSet[num]++
    }

    for {
        if original > 1001 || numSet[original] == 0 {
            break
        }

        numSet[original] = original * 2
        original = original * 2
    }

    return original
}

func main () {

}