func uniformArray(nums1 []int) bool {
    o, e := true, true
    for i := 0; i < len(nums1); i++ {
        d1, d2 := false, false
        if i + 1 < len(nums1){
           if (nums1[i] - nums1[i + 1]) % 2 != 0 {
               d2 = true
           }
        }
        if nums1[i] % 2 != 0 {
            d1 = true
        }
        if !d1 && !d2 {
            o = false
        }
    }
    for i := 0; i < len(nums1); i++ {
        d1, d2 := false, false
        if i + 1 < len(nums1){
           if (nums1[i] - nums1[i + 1]) % 2 == 0 {
               d2 = true
           }
        }
        if nums1[i] % 2 == 0 {
            d1 = true
        }
        if !d1 && !d2 {
            o = false
        }
    }
    return o || e
}
