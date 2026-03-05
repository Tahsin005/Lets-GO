func minOperations(s string) int {
	zeroStateForFirstCandidate := true   // pattern: 0 1 0 1 ...
	zeroStateForSecondCandidate := false // pattern: 1 0 1 0 ...

	count1, count2 := 0, 0

	for _, char := range s {
		if char == '0' {
			if !zeroStateForFirstCandidate {
				count1++
			}
		} else {
			if zeroStateForFirstCandidate {
				count1++
			}
		}

		if char == '0' {
			if !zeroStateForSecondCandidate {
				count2++
			}
		} else {
			if zeroStateForSecondCandidate {
				count2++
			}
		}

		zeroStateForFirstCandidate = !zeroStateForFirstCandidate
		zeroStateForSecondCandidate = !zeroStateForSecondCandidate
	}

	if count1 < count2 {
		return count1
	}
	return count2
}
