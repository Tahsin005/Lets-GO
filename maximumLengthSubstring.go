func maximumLengthSubstring(s string) int {
    count := make(map[byte]int)
    longest := 0
    for l, r := 0, 0; r < len(s); r++ {
        for count[s[r]] == 2 {
            count[s[l]] -= 1
            l++
        }
        count[s[r]] += 1
        longest = max(r - l + 1, longest)
    }
    return longest
}
