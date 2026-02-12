func longestBalanced(s string) int {
    var result int

    r := []rune(s)
    for i := 0; i < len(r); i++ {
        m := make(map[rune]int)
        for j := i; j < len(r); j++ {
            m[r[j]]++
            var freq int
            for _,w := range m {
                if freq == 0 {
                    freq = w
                } else {
                    if freq != w {
                        freq = 0
                        break
                    }
                }
            }
            if result < freq * len(m) {
                result = freq * len(m)
            }
        }
    }
    
    return result
}
