func rearrangeString(s string, x byte, y byte) string {
    m := make(map[byte]int)
    for _,v := range []byte(s) {
        m[v]++
    }
    result := step(y, m[y])
    for v,w := range m {
        if v != y {
            result += step(v, w)
        }
    }
    return result
}

func step(a byte, n int) string {
    var s string
    for i := 0; i < n; i++ {
        s += string(a)
    }
    return s
}
