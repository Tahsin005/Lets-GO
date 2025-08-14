package main

import "fmt"

// https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/?envType=daily-question&envId=2025-08-14

func largestGoodInteger(num string) string {
	var r string

	for i := 0; i < len(num) - 2; i++ {
		if num[i] == num[i + 2] && num[i + 1] == num[i + 2] && r < num[i:i + 3] {
			r = num[i : i + 3]
		}
	}

	return r
}

func main () {

}