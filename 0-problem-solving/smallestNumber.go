func smallestNumber(pattern string) string {
    stack := []rune{}
    var ans string
    pattern += "I"
    for i := 1; i <= len(pattern); i++ {
        stack = append(stack, rune(i + '0'))
        if pattern[i - 1] == 'I' {
            slices.Reverse(stack)
            ans += string(stack)
            stack = stack[:0]
        } 
    }
    return ans
}
