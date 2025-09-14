package main

import "fmt"

// https://leetcode.com/problems/vowel-spellchecker/description/?envType=daily-question&envId=2025-09-14

func spellchecker(wordlist []string, queries []string) []string {
	wordSet, lowerMap, vowelMap := make(map[string]bool), make(map[string]string), make(map[string]string)
	for _, word := range wordlist {
		wordSet[word] = true
		lowerWord := strings.ToLower(word)
		if _, ok := lowerMap[lowerWord]; !ok {
			lowerMap[lowerWord] = word
		}
		vowelWord := replaceVowel(word)
		if _, ok := vowelMap[vowelWord]; !ok {
			vowelMap[vowelWord] = word
		}
	}
	results := make([]string, len(queries))
	for i, query := range queries {
		if wordSet[query] {
			results[i] = query
		} else if word, ok := lowerMap[strings.ToLower(query)]; ok {
			results[i] = word
		} else if word, ok := vowelMap[replaceVowel(query)]; ok {
			results[i] = word
		} else {
			results[i] = ""
		}
	}
	return results
}

func replaceVowel(word string) string {
	vowels := "aeiou"
	mapper := func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return '*'
		}
		return r
	}
	return strings.Map(mapper, strings.ToLower(word))
}

func main () {

}