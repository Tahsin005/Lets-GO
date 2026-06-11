func replaceDigits(s string) string {
    b := []byte(s)
    for i := 1; i < len(s); i += 2 {
        b[i] += b[i - 1] - '0'
    }
    return string(b)
}
