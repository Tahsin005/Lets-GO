func maxDistance(colors []int) int {
	n := len(colors)
	if colors[n - 1] != colors[0] {
		return n - 1
	}

	second := 1
	for second < n && colors[second] == colors[0] {
		second++
	}
	if second == n {
		return n - 1
	}

	maxD := second
	for i := second + 1; i < n; i++ {
		if colors[i] != colors[0] {
			maxD = i
		} else {
			maxD = max(maxD, i - second)
		}
	}

	return maxD
}
