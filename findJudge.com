func findJudge(n int, trust [][]int) int {
    mp := make([]int, n)
    for _, tr := range trust {
        mp[tr[0] - 1]--
        mp[tr[1] - 1]++
    }
    
    for p, t := range mp {
        if t >= n - 1 {
            return p + 1
        }
    }
    return -1
}
