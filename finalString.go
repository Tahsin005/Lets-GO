func finalString(s string) string {
    rn := []rune(s)
    for i := 0; i < len(rn); i++ {
        if rn[i] == 'i' {
            rn = append(reverse(rn[0:i]), rn[i + 1:]...)
            i--
        }
    }
    return string(rn)
}


func reverse(rn []rune) []rune {
    l, r := 0, len(rn) - 1
    for l < r {
        rn[l], rn[r] = rn[r], rn[l] 
        l++
        r--
    }
    return rn
}
