func canMakeSubsequence(s string, t string) bool {
	i, j := 0, 0
	n := len(s)

	for _, ch := range t {
		c := byte(ch)

		a := j
		if j < n && c == s[j] {
			a++
		}

		b := i + 1
		if a > b {
			j = a
		} else {
			j = b
		}

		if i < n && c == s[i] {
			i++
		}
	}

	return j >= n
}
