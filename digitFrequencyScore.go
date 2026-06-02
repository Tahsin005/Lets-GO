func digitFrequencyScore(n int) int {
    var sum int
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}
