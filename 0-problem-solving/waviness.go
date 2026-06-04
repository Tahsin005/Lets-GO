func totalWaviness(num1 int, num2 int) int {
    ans := 0
    for ; num1 <= num2; num1++ {
        ans += waviness(num1)
    }
    return ans
}

func waviness(x int) int {
    ans := 0
    for x > 99 {
        if x % 10 < x % 100 / 10 && x % 100 / 10 > x % 1000 / 100 ||
         x % 10 > x % 100 / 10 && x % 100 / 10 < x % 1000 / 100 {
            ans++
        }
        x /= 10
    }
    return ans
}
