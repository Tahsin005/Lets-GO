func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	seen := make([]int, n+1)
	ans := []int{}
	common := 0

	for i := 0; i < n; i++ {
		seen[A[i]]++
		if seen[A[i]] == 2 {
			common++
		}

		seen[B[i]]++
		if seen[B[i]] == 2 {
			common++
		}

		ans = append(ans, common)
	}

	return ans
}
