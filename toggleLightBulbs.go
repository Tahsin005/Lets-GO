func toggleLightBulbs(bulbs []int) []int {
    state := make([]bool, 101)
    for _, b := range bulbs {
        state[b] = !state[b]
    }
    result := make([]int, 0, 100)
    for i := 1; i <= 100; i++ {
        if state[i] {
            result = append(result, i)
        }
    }
    return result
}
