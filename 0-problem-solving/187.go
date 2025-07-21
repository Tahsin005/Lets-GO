package main

import "fmt"

// https://leetcode.com/problems/delete-characters-to-make-fancy-string/description/?envType=daily-question&envId=2025-07-21

func makeFancyString(s string) string {
	var result strings.Builder
	lastChar := ""
	count := 0

	for i := 0; i < len(s); i++ {
		currentChar := string(s[i])
		if currentChar == lastChar {
			count++
		} else {
			count = 1
			lastChar = currentChar
		}
		if count > 2 {
			continue
		}
		result.WriteString(currentChar)
	}

	return result.String()
}

func main () {

}