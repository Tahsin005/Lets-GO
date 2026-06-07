func consecutiveSetBits(n int) bool {
	binary := strconv.FormatInt(int64(n), 2)
	i, j, count := 0, 1, 0

	for i < j && j < len(binary) {
		if binary[i] == 49 && binary[j] == 49 {
			count++
		}
		i++
		j++
	}

	return count == 1
}
