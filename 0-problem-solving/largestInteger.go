func largestInteger(n int, s int) int {
	if s == 0 {
		return 0
	}

	if s > 9 * n {
		return -1
	}

	ans := ""

	for i := 0; i < n; i++ {
		digit := min(9, s)
		ans += string(rune('0' + digit))
		s -= digit
	}

	res, _ := strconv.Atoi(ans)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
