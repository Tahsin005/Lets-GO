func countDistinctIntegers(nums []int) int {
	m, arr := map[int]bool{}, []int{}
	for i := 0; i < len(nums); i++{
		s, t := strconv.Itoa(nums[i]), ""
		for j := len(s) - 1; j >= 0; j-- {
			t += string(s[j])
		}
		n, _ := strconv.Atoi(t)
		arr = append(arr, n)
	}
	nums = append(nums, arr...)
	for _, v := range nums {
		m[v] = true
	}
	return len(m)
}
