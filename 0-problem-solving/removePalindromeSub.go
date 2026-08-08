func removePalindromeSub(s string) int {
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return 2
        } 
        l++
        r--
    }
    return 1
}
