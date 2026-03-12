func minOperations(s string) int {
	n := len(s)

	sorted := true
	for i := 1; i < n; i++ {
		if s[i] < s[i - 1] {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}

	if n == 2 {
		return -1
	}

	minChar := s[1]
	maxChar := s[1]

	for i := 1; i < n - 1; i++ {
		if s[i] < minChar {
			minChar = s[i]
		}
		if s[i] > maxChar {
			maxChar = s[i]
		}
	}

	if s[0] <= s[n - 1] && (minChar >= s[0] || maxChar <= s[n - 1]) {
		return 1
	}

	if s[n - 1] < minChar && s[0] > maxChar {
		return 3
	}

	return 2
}
