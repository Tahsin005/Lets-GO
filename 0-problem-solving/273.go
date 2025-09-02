package main

import "fmt"

// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/description/?envType=problem-list-v2&envId=n6ww7n9h

func freqAlphabets(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		if i + 2 < len(s) && s[i + 2] == '#' {
			n = 10 * n + s[i + 1] - '0'
			i += 2
		}
		result += string(n + 'a' - 1)
	}
	return result
}

func main () {

}