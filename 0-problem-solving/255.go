package main

import "fmt"

// https://leetcode.com/problems/decode-string/description/

func decodeString(s string) string {
    stringStack := []string{}
	numberStack := []int{}
	current := ""
	num := 0

	for _, c := range s {
		if unicode.IsDigit(c) {
			num = num * 10 + int(c - '0')
		} else if c == '[' {
			numberStack = append(numberStack, num)
			stringStack = append(stringStack, current)
			num = 0
			current = ""
		} else if c == ']' {
			k := numberStack[len(numberStack) - 1]
			numberStack = numberStack[:len(numberStack) - 1]

			prev := stringStack[len(stringStack) - 1]
			stringStack = stringStack[:len(stringStack) - 1]

			expanded := prev
			for i := 0; i < k; i++ {
				expanded += current
			}
			current = expanded
		} else {
			current += string(c)
		}
	}

	return current
}

func main () {

}