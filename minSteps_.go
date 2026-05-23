func minSteps(s string, t string) int {
    freq := [26]int{}
    for i := range s {
        freq[s[i] - 'a']++
        freq[t[i] - 'a']--
    }
    
    ans := 0
    for _, n := range freq {
        if n > 0 {
            ans += n
        }
    }
    return ans
}
