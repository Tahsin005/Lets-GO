func minimumPushes(word string) int {
    m := make(map[rune]bool)
    for _,v := range word {
        m[v] = true
    }
    n := len(m)
    var result int
    for i := 1; n > 0; i++ {
        result += i * min(8, n)
        n -= 8
    }
    return result
}
