func scoreValidator(events []string) []int {
	score, counter := 0, 0
	for i := 0; i < len(events); i++ {
		if counter == 10 {
			break
		}
		switch events[i] {
		case "W":
			counter++
		case "WD":
			score++
		case "NB":
			score++
		default:
			n, _ := strconv.Atoi(events[i])
			score += n
		}
	}
	return []int{score, counter}
}
