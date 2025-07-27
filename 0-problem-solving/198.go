package main

import "fmt"

// https://leetcode.com/problems/permutation-difference-between-two-strings/description/?envType=problem-list-v2&envId=hash-table

func findPermutationDifference(s string, t string) int {
	abs := func(x int) int {
		return max(x, -x)
	}
	m := map[byte]int{}
	res := 0
	for i := range s {
		m[s[i]] = i
	}
	for i := range t {
		res += abs(m[t[i]] - i)
	}
	return res
}

func main () {

}