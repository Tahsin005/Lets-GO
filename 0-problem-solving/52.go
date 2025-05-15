package main

import "fmt"

// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/description/?envType=daily-question&envId=2025-05-15

func getLongestSubsequence(words []string, groups []int) []string {
    res := []string{words[0]}
    prev := groups[0]

    for i := 1; i < len(words); i++ {
        if groups[i] != prev {
            res = append(res, words[i])
            prev = groups[i]
        }
    }

    return res
}

func main () {

}