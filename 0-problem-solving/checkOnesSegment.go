func checkOnesSegment(s string) bool {
    seenZero := false

    for _, c := range s {
        if c == '0' {
            seenZero = true
        } else if seenZero {
            return false
        }
    }

    return true
}
