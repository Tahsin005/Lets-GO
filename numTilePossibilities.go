func numTilePossibilities(tiles string) int {
	tiles_b := []byte(tiles)
	slices.Sort(tiles_b)

	segments := []int{}
	for l := 0; l < len(tiles_b); {
		r := l + 1
		for r < len(tiles_b) && tiles_b[r] == tiles_b[l] {
			r++
		}
		segments = append(segments, r - l)
		l = r
	}

    fmt.Println(segments)

	return backtrack(segments)
}

func backtrack(segments []int) int {
	answer := 0
	for i := range segments {
		if segments[i] > 0 {
			segments[i]--
			answer += 1 + backtrack(segments)
			segments[i]++
		}
	}
	return answer
}
