func countCommas(n int) int {
	if n - 1000 < 0 {
		return 0
	}
	return n - 1000 + 1
}
