package main

import "fmt"

// https://leetcode.com/problems/three-consecutive-odds/description/?envType=daily-question&envId=2025-05-11

func threeConsecutiveOdds(arr []int) bool {
    for i := 0; i <= len(arr) - 3; i++ {
        if arr[i] % 2 == 1 && arr[i + 1] % 2 == 1 && arr[i + 2] % 2 == 1 {
            return true
        }
    }
    return false
}

func main () {

}