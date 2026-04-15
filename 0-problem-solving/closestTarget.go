func closestTarget(words []string, target string, start int) int {
	n := len(words)
	for i := range n / 2 + 1 {
        l, r := (start - i + n) % n, (i + start) % n
		if words[l] == target || words[r] == target {
			return i
		}
	}
	return -1
}
