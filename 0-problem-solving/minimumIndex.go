func minimumIndex(capacity []int, itemSize int) int {
	idx, min := -1, itemSize
	for i := 0; i < len(capacity); i++ {
		subs := capacity[i] - itemSize
		if idx == -1 && subs >= 0 {
			idx = i
			min = subs
		}
		if subs < min && subs >= 0 {
			idx = i
			min = subs
		}
	}
	return idx
}
