func shortestBeautifulSubstring(s string, k int) string {
	mini := 101
	ans := ""

	for l, r, sum := 0, 0, 0; r < len(s); r++ {
		if s[r] == '1' {
			sum++
		}
		for sum == k {
			curr := r - l + 1
			sub := s[l : r + 1]
			if curr < mini {
				mini = curr
				ans = sub
			} else if curr == mini && sub < ans {
				ans = sub
			}

			if s[l] == '1' {
				sum--
			}
			l++
		}
	}
	return ans
}
