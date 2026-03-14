func getHappyString(n int, k int) string {
	happyStrings := []string{}

	var generate func(current string)
	generate = func(current string) {
		if len(current) == n {
			happyStrings = append(happyStrings, current)
			return
		}

		for c := 'a'; c <= 'c'; c++ {
			if len(current) > 0 && rune(current[len(current) - 1]) == c {
				continue
			}

			generate(current + string(c))
		}
	}

	generate("")

	if len(happyStrings) < k {
		return ""
	}

	return happyStrings[k-1]
}
