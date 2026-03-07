func minFlips(s string) int {
	if len(s) & 1 == 0 {
		return minFlipsEvenLen(s)
	} else {
		return minFlipsOddLen(s)
	}
}

func minFlipsEvenLen(s string) int {
	p1 := 0
	for i := range s {
		p1 += (int(s[i]) ^ i) & 1
	}
	return min(p1, len(s) - p1)
}

func minFlipsOddLen(s string) int {
	minP3 := -1
	maxP3 := -1
	p1 := 0
	for i := range s {
		p1 += (int(s[i]) ^ i) & 1
		minP3 = min(minP3, i - p1 * 2)
		maxP3 = max(maxP3, i - p1 * 2)
	}
	minP3 += p1 + 1
	maxP3 += p1 + 1
	return min(minP3, len(s) - maxP3)
}
