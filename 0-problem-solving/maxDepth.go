func maxDepth(s string) int {
    count, max := 0, 0
    for _, r := range s {
        switch r {
            case '(':
            count++
            case ')':
            count--
        }
        if count > max {
            max = count
        }
    }
    return max
}
