package main

import "fmt"

// https://leetcode.com/problems/substring-matching-pattern/description/

func hasMatch(s string, p string) bool {
	idx := strings.IndexRune(p, '*')
	prefix := p[:idx]
	suffix := p[idx + 1:]

	prefixIdx := strings.Index(s, prefix)
	if prefixIdx == -1 {
		return false
	}

	suffixIdx := strings.LastIndex(s, suffix)
	if suffixIdx == -1 {
		return false
	}

	return suffixIdx >= prefixIdx + len(prefix)
}

func main () {

}