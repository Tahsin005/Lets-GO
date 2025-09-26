package main

import "fmt"

// https://leetcode.com/problems/valid-triangle-number/description/?envType=daily-question&envId=2025-09-26

func triangleNumber(sides []int) int {
	sort.Ints(sides)
	totalTriangles := 0

	for longest := len(sides) - 1; longest >= 2; longest-- {
		left, right := 0, longest - 1
		for left < right {
			if sides[left] + sides[right] > sides[longest] {
				totalTriangles += (right - left)
				right--
			} else {
				left++
			}
		}
	}

	return totalTriangles
}

func main () {

}