func elevatorRequests(n int, requests []int) int {
    total := 0
    curr := 0
    
    for _, floor := range requests {
        if curr > floor {
            total += curr - floor
        } else {
            total += floor - curr
        }
        curr = floor
    }
    
    return total
}
