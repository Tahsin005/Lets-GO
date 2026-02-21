func countPrimeSetBits(left int, right int) int {
    res := 0
    for i := left; i <= right; i++ {
        s := setB(i)
        if isPrime(s) {
            res++
        }
    }
    return res
}
func setB(n int)int{
    res := 0
    for n != 0 {
        n &= (n - 1)
        res++
    }
    return res
}
func isPrime(n int)bool{
    if n < 2 {
        return false
    }
    for i := 2; i <= sqrt(n); i++ { 
        if n % i == 0 {
            return false
        }
    }
    return true
}
func sqrt(x int) int {
    if x == 1 {
        return 1
    }
    y := x / 2
    for y * y > x {
        y = (y + x / y) >> 1
    }
    return y
}
