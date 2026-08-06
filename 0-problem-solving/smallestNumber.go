func smallestNumber(n int, t int) int {
    for mult(n) % t != 0 {
        n++
    }
    return n
}

func mult(n int) int {
    mult := 1
    for n > 0 {
        mult *= n % 10 
        n /= 10
    }
    return mult
}
