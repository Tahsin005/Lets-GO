func mirrorFrequency(s string) (r int) {

	appear := map[string]bool{}

	for i := 0; i < len(s); i++ {
		ori, mirror := string(s[i]), ""
		if s[i] >= '0' && s[i] <= '9' {
			mirror = string('9' - (rune(s[i]) - '0'))
		} else if s[i] >= 'a' && s[i] <= 'z' {
			mirror = string('z' - (rune(s[i]) - 'a'))
		}

		if appear[ori] || appear[mirror] {
			continue
		}

		x, y := 0, 0
		for j := 0; j < len(s); j++ {
			if string(s[j]) == mirror {
				y++
			}
			if string(s[j]) == ori {
				x++
			}
		}

		if x > 0 {
			appear[string(s[i])] = true
		}
		if y > 0 {
			appear[mirror] = true
		}
		r += abs(x - y)
	}
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
