package main

import "fmt"

// https://leetcode.com/problems/bitwise-or-of-even-numbers-in-an-array/

func evenNumberBitwiseORs(nums []int) int {
    count := 0

    for _, num := range nums {
        if num % 2 == 0 {
            count |= num
        }
    }

    return count
}

func main () {

}