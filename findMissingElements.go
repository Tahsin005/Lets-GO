func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}
func min (a, b int) int{
    if a < b {
        return a
    }
    return b
}
func findMissingElements(nums []int) []int {
    maximum := -1 << 32
    minimum := 1 << 32
    items := map[int]bool{}
    results := []int{}
    for _, num := range nums {
        items[num] = true
        maximum = max(maximum, num)
        minimum = min(minimum, num)
    }
    for i := minimum; i < maximum; i++ {
        if _, ok := items[i]; !ok {
            results = append(results, i)
        }
    }
    return results
}
