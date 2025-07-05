package main

import "fmt"

// https://leetcode.com/problems/find-lucky-integer-in-an-array/description/?envType=daily-question&envId=2025-07-05

func findLucky(arr []int) int {
    var max int
    for i := 0; i < len(arr); i++ {
        if max < arr[i] {
            max = arr[i]
        }
    }
    frq := make([]int, max + 1)
    for i := 0; i < len(arr); i++ {
        frq[arr[i]]++
    }
    for i := max; i > 0; i--  {
        if i == frq[i]  {
            return i
        }
    }
    return -1
}

func main () {

}