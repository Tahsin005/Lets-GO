func checkStrings(s1 string, s2 string) bool {
    even := [26]int{}
    odd := [26]int{}

    for i := range len(s1) {
        if i % 2 == 0 {
            even[s1[i] - 'a']++
            even[s2[i] - 'a']--
        } else {
            odd[s1[i] - 'a']++
            odd[s2[i] - 'a']--
        }
    }

    return even == [26]int{} && odd == [26]int{}
}
