func freqAlphabets(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		if i + 2 < len(s) && s[i + 2] == '#' {
			n = 10 * n + s[i + 1] - '0'
			i += 2
		}
		result += string(n + 'a' - 1)
	}
	return result
}