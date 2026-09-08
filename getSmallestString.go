func getSmallestString(s string) string {
    n := len(s)
	charArray := []rune(s)

	for i := 0; i < n - 1; i++ {
		a := charArray[i] - '0'
		b := charArray[i + 1] - '0'
		if a % 2 == b % 2 && a > b {
			charArray[i], charArray[i + 1] = charArray[i + 1], charArray[i]
			break
		}
	}

	return string(charArray)
}
