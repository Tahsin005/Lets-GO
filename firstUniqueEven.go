func firstUniqueEven(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if v % 2 == 0 {
			m[v]++
		}
	}
	for _, v := range nums {
		if m[v] == 1 {
			return v
		}
	}
	return -1
}
