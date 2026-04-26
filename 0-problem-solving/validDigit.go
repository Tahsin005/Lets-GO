func validDigit(n int, x int) bool {
    s := strconv.Itoa(n)
    digit := strconv.Itoa(x)
    
    return strings.Contains(s, digit) && s[0] != digit[0]
}
