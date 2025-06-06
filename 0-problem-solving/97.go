package main

import "fmt"

// https://leetcode.com/problems/split-strings-by-separator/description/?envType=problem-list-v2&envId=array

func splitWordsBySeparator(words []string, separator byte) []string {
	s := string(separator)
	ansTmp, ans := []string{}, []string{}
	for _, w := range words {
		ansTmp = append(ansTmp, strings.Split(w, s)...)
	}

    fmt.Println(ansTmp)
	for _, a := range ansTmp {
		if a != "" {
			ans = append(ans, a)
		}
	}
	return ans
}


func main () {

}