func minPenalty(period int, lights []int, arrivalTime []int) int {
    maxi := lights[0]
    for _, v := range lights {
        if v > maxi {
            maxi = v
        }
    }
    res := 0
    
    for _, t := range arrivalTime {
        r := t % period
        if r >= maxi {
            penalty := period - r
            if penalty > res {
                res = penalty
            }
        }
    }
    
    return res
}
