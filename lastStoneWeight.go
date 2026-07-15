func lastStoneWeight(stones []int) int {
    for {
        idx1, st1 := max(stones)
        stones[idx1] = 0
        idx2, st2 := max(stones)
        stones[idx2] = 0
        stones = append(stones, abs(st1 - st2))
        if st1 == 0 || st2 == 0 {
            break
        }
    }
    return slices.Max(stones)
}

func abs(x int) int {
    if x < 0 {
        x = -x
    }
    return x
}

func max(stones []int) (int, int) {
    max, idx := 0, 0
    for i, s := range stones {
        if s > max {
            max = s
            idx = i
        }
    }
    return idx, max
}
