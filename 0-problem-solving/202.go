package main

import "fmt"

// https://leetcode.com/problems/decode-the-message/description/?envType=problem-list-v2&envId=hash-table

func decodeMessage(key string, message string) string {
	var sw = 'a'
	var alph = make(map[rune]rune, 26)

	alph[' '] = ' '
	for _, r := range key {
		if _, exists := alph[r]; !exists {
			alph[r] = sw
			sw++
		}
	}

	var r = make([]rune, len(message))
	for i := 0; i < len(message); i++ {
		r[i] = alph[rune(message[i])]
	}

	return string(r)
}

func main () {

}