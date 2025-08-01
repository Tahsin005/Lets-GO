package main

import "fmt"

// https://leetcode.com/problems/count-elements-with-maximum-frequency/description/?envType=problem-list-v2&envId=hash-table

func maxFrequencyElements(nums []int) int {
    mp := make(map[int]int)
    max := 0
    for _, v := range nums {
        mp[v]++
        if mp[v] > max {
            max = mp[v]
        }
    }

    answer := 0
    for _, v := range mp {
        if v == max {
            answer += v
        }
    }
    return answer
}

func main () {

}