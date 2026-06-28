func maximumElementAfterDecrementingAndRearranging(a []int) int {
	slices.Sort(a)
	p := 1
	for _, v := range a {
		if v > p {
			v = p
		}
		p = v + 1
	}
	return p - 1
}
