func isAdjacentDiffAtMostTwo(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
		x, _ := strconv.Atoi(string(s[i]))
		y, _ := strconv.Atoi(string(s[i + 1]))
		if abs(x - y) > 2 {
			return false
		}

	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
