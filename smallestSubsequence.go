func smallestSubsequence(s string) string {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch - 'a']++
	}

	inStack := make([]bool, 26)

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]

		freq[ch - 'a']--

		if inStack[ch - 'a'] {
			continue
		}

		for len(stack) > 0 && stack[len(stack) - 1] > ch && freq[stack[len(stack) - 1] - 'a'] > 0 {
			inStack[stack[len(stack)-1] - 'a'] = false
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, ch)
		inStack[ch - 'a'] = true
	}

	return string(stack)
}
