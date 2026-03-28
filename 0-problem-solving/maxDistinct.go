func maxDistinct(s string) int {
    seen := [26]bool{}
    count := 0

    for i := 0; i < len(s); i++ {
        idx := s[i] - 'a'
        if !seen[idx] {
            seen[idx] = true
            count++
            if count == 26 {
                return 26
            }
        }
    }
    return count
}
