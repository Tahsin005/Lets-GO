package main

import "fmt"

// https://leetcode.com/problems/lexicographical-numbers/description/?envType=daily-question&envId=2025-06-08

func lexicalOrder(n int) []int {
    res := []int{}
    curr := 1
    for i := 0; i < n; i++ {
        res = append(res, curr)
        if curr * 10 <= n {
            curr *= 10
        } else {
            for curr % 10 == 9 || curr >= n {
                curr /= 10
            }
            curr++
        }
    }
    return res
}

func main () {

}