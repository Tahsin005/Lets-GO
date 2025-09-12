package main

import "fmt"

// https://leetcode.com/problems/split-a-string-in-balanced-strings/

func balancedStringSplit(s string) int {
	var stack int
	var out int
	for i := range s {
		if s[i] == 'R' {
			stack++
		} else {
			stack--
		}
		if stack == 0 {
			out++
		}
	}
	return out
}

func main () {

}