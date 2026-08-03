func countValidPrefixes(s string) int {
    var n0, n1, result int
    for _,v := range s {
        if v == '0' {
            n0++
        } else {
            n1++
        }
        if n0 + 1 ==n1 || n1 + 1 == n0 || n0 == n1 {
            result++
        }
    }
    return result
}
