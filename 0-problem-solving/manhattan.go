func maxDistance(moves string) int {
	x, y, unknown := 0, 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			x--
		case 'D':
			y--
		case 'R':
			x++
		case 'U':
			y++
		default:
			unknown++
		}
	}
	return abs(x) + abs(y) + unknown
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
