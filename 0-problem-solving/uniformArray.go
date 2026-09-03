func uniformArray(nums []int) bool {
    minny := int(1e9)
    hasOdd := false
    
    for _, num := range nums {
        minny = min(minny, num)
        if num % 2 != 0 {
            hasOdd = true
        }
    }

    if !hasOdd {
        return true
    }

    return minny % 2 != 0
}
