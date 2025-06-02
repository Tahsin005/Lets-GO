package main

import "fmt"

// https://leetcode.com/problems/find-maximum-number-of-string-pairs/?envType=problem-list-v2&envId=array

func maximumNumberOfStringPairs(words []string) int {
	st := make(map[string]bool)
	res := 0
	for _, i := range words {
		s := i
		i = reverseString(i)
		if !st[i] {
			st[s] = true
		} else {
			res++
		}
	}
	return res
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main () {

}