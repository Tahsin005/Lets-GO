func numberOfMatches(n int) int {
    sumMatch := 0
    for n > 0 {
        if n % 2 != 0 {
            sumMatch++
            n--
        }
        sumMatch += n / 2
        n /= 2
    }
    return sumMatch - 1
}
