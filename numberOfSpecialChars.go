func numberOfSpecialChars(word string) int {
    mp := make(map[rune]bool)
    for _, r := range word {
        mp[r] = true
    }
    count := 0
    for k, _ := range mp {
        if mp[k] && mp[k - 32] {
            count++
        }
    }
    return count
}
