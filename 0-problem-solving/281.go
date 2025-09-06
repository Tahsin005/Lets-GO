package main

import "fmt"

// https://leetcode.com/problems/count-items-matching-a-rule/description/?envType=problem-list-v2&envId=n6ww7n9h

func countMatches(items [][]string, ruleKey string, ruleValue string) (r int) {
	for i := 0; i < len(items); i++ {
		n := map[string]int {
			"type": 0,
			"color": 1,
			"name": 2,
		}
		for j := 0; j < len(items[i]); j++ {
			if j == n[ruleKey] && items[i][j] == ruleValue {
				r++
			}
		}
	}
	return
}

func main () {

}