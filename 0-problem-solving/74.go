package main

import "fmt"

// https://leetcode.com/problems/find-words-containing-character/description/?envType=daily-question&envId=2025-05-26

func findWordsContaining(words []string, x byte) []int {
	res := []int{}
	for index, word := range words {
		for _, c := range word {
			if c == rune(x) {
				res = append(res, index)
				break
			}
		}
	}
	return res
}

func main () {

}