func checkDivisibility(n int) bool {
    product := 1
    sum := 0
    num := n
    for num > 0 {
        rem := num % 10
        num = num / 10
        sum += rem
        product *= rem
    }
    if n % (sum + product) == 0 {
        return true
    }
    return false
}
