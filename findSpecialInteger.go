func findSpecialInteger(arr []int) int {
    mp := make(map[int]int)
    for _, n := range arr {
        mp[n]++
        if mp[n] > len(arr) / 4 {
            return n
        }
    }
    return 0
}
