package main

import "fmt"

// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/description/?envType=daily-question&envId=2025-10-08

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := []int{}
	for i := 0; i < len(spells); i++ {
		l, r := 0, len(potions) - 1
		for l <= r {
			m := l + (r - l) / 2
			if int64(spells[i] * potions[m]) < success {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		res = append(res, len(potions) - l)
	}
	return res
}

func main () {

}