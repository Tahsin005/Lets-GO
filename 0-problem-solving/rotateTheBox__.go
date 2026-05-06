func rotateTheBox(boxGrid [][]byte) [][]byte {
    m := len(boxGrid)
    n := len(boxGrid[0])
    
    // Create the rotated result grid (n x m)
    result := make([][]byte, n)
    for i := range result {
        result[i] = make([]byte, m)
    }
    
    // 1. Simulate gravity row by row (falling to the right)
    for row := 0; row < len(boxGrid); row++ {
        empty := len(boxGrid[row]) - 1 // Tracks the lowest available 'floor'
        
        // Scan from right to left
        for i := len(boxGrid[row]) - 1; i >= 0; i-- {
            
            if boxGrid[row][i] == '#' {
                // If it's a stone, move it to the furthest empty slot
                boxGrid[row][i] = '.'
                boxGrid[row][empty] = '#'
                empty-- // The floor rises by one
                
            } else if boxGrid[row][i] == '*' {
                // If it's an obstacle, reset the floor to just above the obstacle
                empty = i - 1
            }
        }
    }
    
    // 2. Rotate the settled box 90 degrees clockwise
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            result[j][m-1-i] = boxGrid[i][j]
        }
    }
    
    return result
}
