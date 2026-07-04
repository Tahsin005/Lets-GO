func hammingDistance(x int, y int) int {
    z := x ^ y
    count := 0
    for z != 0 {
        if z & 1 == 1 {
            count++
        }
        z >>= 1
    }
    return count
}
