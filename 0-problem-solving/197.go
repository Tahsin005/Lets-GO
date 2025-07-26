package main

import "fmt"

// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/description/?envType=problem-list-v2&envId=hash-table

func maxFreqSum(s string) int {
	freq := map[byte]int{}
	for i := range s {
		freq[s[i]]++
	}
	maxVowelFreq := 0
	for _, ch := range "aeuio" {
		maxVowelFreq = max(maxVowelFreq, freq[byte(ch)])
		freq[byte(ch)] = 0
	}
	maxConsFreq := 0
	for _, val := range freq {
		maxConsFreq = max(maxConsFreq, val)
	}
	return maxConsFreq + maxVowelFreq
}

func main () {

}