package main

import "fmt"

// https://leetcode.com/problems/count-largest-group/description/?envType=daily-question&envId=2025-04-27

func countLargestGroup(n int) int {
    digitSum := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        digitSum[i] = digitSum[i / 10] + i % 10
    }

    freq := make(map[int]int, n)
    maxFreq := 0
    for i := 1; i <= n; i++ {
        s := digitSum[i]
        freq[s]++
        if freq[s] > maxFreq {
            maxFreq = freq[s]
        }
    }
    
    result := 0
    for _, count := range freq {
        if count == maxFreq {
            result++
        }
    }
    return result
}

func main () {
	
}