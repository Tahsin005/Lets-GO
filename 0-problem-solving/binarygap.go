func binaryGap(n int) int {
    maxgap := 0
    last := -1
    for i := 0; n > 0; i++ {       
        if n & 1 == 1 {
            if last != -1 && i - last > maxgap  {
                maxgap = i - last
            }
            last = i
        }        
        n >>= 1
    }
    return maxgap
}