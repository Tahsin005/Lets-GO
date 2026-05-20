func isPerfect(val int, k int,l int, r int) bool {
    p := int(math.Pow(float64(val), float64(k)))
    if p >= l && p <= r {
        return true
    }
    return false
}

func countKthRoots(l int, r int, k int) int {
    if k == 1  {
        return r - l + 1
    }
    
    count := 0
    for i := 1; int(math.Pow(float64(i), float64(k))) <= r; i++ {
        if isPerfect(i, k, l, r) {
            count++
        }
    }

    if l == 0 {
        count++
    }
    return count
}
