func removeDuplicates(s string) string {
    res := make([]rune, 0, len(s))
    for _, r := range s {
        if len(res) > 0 && r == res[len(res) - 1] {
            res = res[:len(res) - 1]
            continue
        }
        res = append(res, r)
    }
    return string(res)
}