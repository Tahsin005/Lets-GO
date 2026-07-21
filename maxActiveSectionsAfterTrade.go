func maxActiveSectionsAfterTrade(s string) int {
    n := len(s)
    ones := 0
    i := 0
    for i < n && s[i] != '0' {
        ones++
        i++
    }

    last := byte('0')
    chain := 0
    runs := []int{}
    
    for ; i < n; i++ {
        if s[i] == '1' {
            ones++
        }
        if s[i] == last {
            chain++
        } else {
            runs = append(runs, chain)
            chain = 1
            last = s[i]
        }
    }
    runs = append(runs, chain)

    out := ones

    for i := 0; i < len(runs) - 2; i += 2 {
        out = max(out, ones + runs[i] + runs[i + 2])
    }

    return out
}
