package main

import "fmt"

// https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/description/?envType=problem-list-v2&envId=string

func validStrings(n int) (r []string) {
	if n == 1 {
		return []string{"0", "1"}
	}

	x := 0
	for 0 < 1 {
		s := strconv.FormatInt(int64(x), 2)
		if len(s) > n{
			break
		}
		if !strings.Contains(s, "00") && strings.Contains(s, "1")  {
			if len(s) == n || len(s) == n - 1 {
				if len(s) == n - 1 {
				    r = append(r, "0" + s)
			    } else {
				    r = append(r, s)
			    }
			}
		}
		x++
	}

	return
}

func main () {

}