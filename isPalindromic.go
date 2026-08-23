func isPalindromic(s string) bool {
    var sb strings.Builder
    for i := 0; i < len(s); i++ {
        b := s[i]
        for j := 7; j >= 0; j-- {
            if (b>>uint(j))&1 == 1 {
                sb.WriteByte('1')
            } else {
                sb.WriteByte('0')
            }
        }
    }

    t := sb.String()
    for i, j := 0, len(t) - 1; i < j; i, j = i + 1, j - 1 {
        if t[i] != t[j] {
            return false
        }
    }
    return true
}
