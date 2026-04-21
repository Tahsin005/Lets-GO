func firstStableIndex(n []int, k int) int {
    if len(n) == 1 {
        return 0
    }

    l, r := make([]int, len(n)), make([]int, len(n))
    l[0], r[len(r) - 1] = n[0], n[len(n) - 1]
    for i := 1; i < len(n); i++ {
        l[i] = max(n[i], l[i - 1])
        r[len(n) - 1 - i] = min(n[len(n) - 1 - i], r[len(n) - i])
    }

    for i := 0; i < len(n); i++ {
        d := l[i] - r[i]
        if d <= k {
            return i
        }
    }

    return -1
}
