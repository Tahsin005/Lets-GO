package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://codeforces.com/problemset/problem/1606/A

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()

		if len(s) > 0 {
			lastChar := string(s[len(s)-1])
			rest := s[1:]
			fmt.Println(lastChar + rest)
		}
	}
}
