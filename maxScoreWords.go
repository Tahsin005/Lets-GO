func maxScoreWords(words []string, letters []byte, score []int) int {
	lettersCount := [26]int{}
	for i := 0; i < len(letters); i++ {
		lettersCount[int(letters[i] - 'a')]++
	}

	var recurse func(int, [26]int) int
	recurse = func(idx int, count [26]int) int {
		if idx == len(words) {
			return 0
		}
		skip, take := recurse(idx + 1, count), 0
		if newPairs, gains, ok := canTake(count, words[idx], score); ok {
			take = gains + recurse(idx + 1, newPairs)
		}
		return max(skip, take)
	}
	return recurse(0, lettersCount)
}

func canTake(count [26]int, word string, score []int) ([26]int, int, bool) {
	gains := 0
	for i := range word {
		letterIdx := int(word[i] - 'a')
		if count[letterIdx] == 0 {
			return [26]int{}, 0, false
		}
		count[letterIdx]--
		gains += score[letterIdx]
	}
	return count, gains, true
}
