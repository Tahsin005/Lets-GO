func bitwiseComplement(n int) int {
    if n == 0 {
        return 1
    }

    bits := int(math.Log2(float64(n))) + 1
    mask := (1<<bits) - 1

    return n ^ mask
}
