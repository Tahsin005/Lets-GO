package main

import "fmt"

// https://leetcode.com/problems/count-prefixes-of-a-given-string/description/?envType=problem-list-v2&envId=array

func countPrefixes(words []string, s string) int {
	var count int

	for _, v := range words {
		if strings.HasPrefix(s, v) {
			count++
		}
	}

	return count
}

func main () {

}