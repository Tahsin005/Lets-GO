func checkGoodInteger(n int) bool {
    var sum int
    for n > 0 {
        a := n % 10
        sum += a * a - a
        if sum >= 50 {
            return true
        }
        n /= 10
    }
    return false
}
