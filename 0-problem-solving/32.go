package main

import "fmt"

// https://leetcode.com/problems/shortest-distance-to-a-character/description/?envType=problem-list-v2&envId=array

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func shortestToChar(s string, c byte) []int {
    walkLeftToRight, currPos := make([]int, len(s)), -len(s)

    for i := range s {
        if s[i] == c {
            currPos = i
        } else {
            walkLeftToRight[i] = i - currPos
        }
    }

    for r := len(s) - 1; r >= 0; r-- {
        if s[r] == c {
            currPos = r
        } else {
            walkLeftToRight[r] = minNum(walkLeftToRight[r], int(math.Abs(float64(currPos) - float64(r))))
        }
    }

    return walkLeftToRight
}

func main () {

}