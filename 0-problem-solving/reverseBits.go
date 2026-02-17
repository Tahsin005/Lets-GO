func reverseBits(num int) int {
    var result int
    for i := 0; i < 32; i++ {
        b := num & 1
        result |= b << (31 - i)
        num >>= 1
    }
    return result
}
