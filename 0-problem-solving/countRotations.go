func countRotations(s string, k int) int {
    n := len(s)

    total := 0

    for i := 0; i < n - 1; i++ {
        if s[i] == s[i + 1] {
            total++
        }
    }

    if s[n - 1] == s[0] {
        total++;
    }

    if k == total {
        return n - total
    }

    if k == total - 1 {
        return total
    }

    return 0
}
