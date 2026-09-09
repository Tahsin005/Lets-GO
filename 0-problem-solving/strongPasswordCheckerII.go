func strongPasswordCheckerII(password string) bool {
    if len(password) < 8 {
        return false
    }
    for i := 0; i < len(password) - 1; i++ {
        if password[i] == password[i + 1] {
            return false
        }
    }
    var a, b, c, d int
    for _, r := range password {
        if unicode.IsUpper(r) {
            a++
        }
        if unicode.IsLower(r) {
            b++
        }
        if strings.Contains("!@#$%^&*()-+", string(r)) {
            c++
        }
        if unicode.IsNumber(r) {
            d++
        }
    }
    return a > 0 && b > 0 && c > 0 && d > 0 
}
