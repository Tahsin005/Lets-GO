package main

import "fmt"

// https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/description/?envType=problem-list-v2&envId=array

func isAcronym(words []string, s string) bool {
	if len(s) != len(words) {
		return false
	}

	for i := range s {
		if s[i] != words[i][0] {
			return false
		}
	}
	return true
}

func main () {

}