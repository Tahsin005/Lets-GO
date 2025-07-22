package main

import "fmt"

// https://leetcode.com/problems/find-indices-of-stable-mountains/description/?envType=problem-list-v2&envId=array

func stableMountains(height []int, threshold int) []int {
    n := len(height)
    var ans []int
    for i := 1; i < n; i++ {
        if height[i - 1] > threshold {
            ans = append(ans, i)
        }
    }

    return ans
}

func main () {

}