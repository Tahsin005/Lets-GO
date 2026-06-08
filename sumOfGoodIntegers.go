func sumOfGoodIntegers(n int, k int) (r int) {
	x := max(1, n - k)
	for i := x; i <= n + k; i++ {
		if abs(n - i) <= k && n & i == 0 {
			r += i
		}
	}
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
