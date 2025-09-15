package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-words-you-can-type/description/?envType=daily-question&envId=2025-09-15

func canBeTypedWords(text string, brokenLetters string) int {
	var broken [26]int
	for i := range brokenLetters {
		broken[brokenLetters[i] - 'a'] = 1
	}

	result := 0
	i := 0
	n := len(text)
	for i < n {
		valid := 1
		for i < n && text[i] != ' ' {
			valid &^= broken[text[i] - 'a']
			i++
		}
		result += valid
		i++
	}
	return result
}

func main () {

}