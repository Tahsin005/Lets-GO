package main

import "fmt"

// https://leetcode.com/problems/truncate-sentence/description/?envType=problem-list-v2&envId=array

func truncateSentence(s string, k int) string {
	res := ""
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if count == k - 1 {
				res = s[:i]
				break
			}

			count++
		}
	}

	if len(res) == 0 {
		return s
	}

	return res
}

func main () {

}