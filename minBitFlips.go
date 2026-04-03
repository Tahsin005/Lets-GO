func minBitFlips(start int, goal int) int {
    flipCount := 0
    mask := 1

    for i := 0; i < 32; i++ {
        startBit := (start & mask) != 0
        goalBit := (goal & mask) != 0

        if startBit != goalBit {
            flipCount++
        }

        mask <<= 1
    }

    return flipCount
}
