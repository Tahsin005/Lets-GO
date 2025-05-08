package main

import "fmt"

// https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/?envType=daily-question&envId=2025-05-08

func numEquivDominoPairs(dominoes [][]int) int {
    count := make([]int, 100)
    res := 0

    for _, d := range dominoes {
        val := d[0] * 10 + d[1]
        if d[0] > d[1] {
            val = d[1] * 10 + d[0]
        }
        res += count[val]
        count[val]++
    }

    return res
}

func main () {

}