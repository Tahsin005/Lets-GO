func hasAllCodes(s string, k int) bool {
    actual := 1 << k

    if len(s) - k + 1 < actual {
        return false
    }

    set := make(map[string]bool)

    for i := 0; i <= len(s) - k; i++ {
        set[s[i:i + k]] = true
    }

    return len(set) == actual
}
