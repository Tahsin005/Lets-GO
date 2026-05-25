func limitOccurrences(nums []int, k int) (r []int) {
	m := map[int]int{}
	for _, n := range nums {
		if m[n] < k {
			r = append(r, n)
			m[n]++
		}
	}
	return
}
