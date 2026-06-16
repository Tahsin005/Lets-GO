func processStr(s string) string {
    res := []rune{}

    for _, c := range s {
        switch c {
        case '*':
            if len(res) > 0 {
                res = res[:len(res) - 1]
            }
        case '#':
            res = append(res, res...)
        case '%':
            for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
                res[i], res[j] = res[j], res[i]
            }
        default:
            res = append(res, c)
        }
    }

    return string(res)
}
