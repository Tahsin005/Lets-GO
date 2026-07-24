func canReach(start []int, target []int) bool {
    manhattanDistance := abs(start[0] - target[0]) + abs(start[1] - target[1])
    return manhattanDistance % 2 == 0
}

func abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}
