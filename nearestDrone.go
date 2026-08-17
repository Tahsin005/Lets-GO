func nearestDrone(drones [][]int, target []int) int {
    dx, dy := target[0], target[1]
    idx := -1
    best := math.MaxInt32
    
    for i, drone := range drones {
        x, y, r := drone[0], drone[1], drone[2]
        dist := abs(x - dx) + abs(y - dy)
        
        if dist <= r && dist < best {
            best = dist
            idx = i
        }
    }
    
    return idx
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
