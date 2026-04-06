type coord struct {
    x, y int
}

func robotSim(commands []int, obstacles [][]int) int {
    obsMap := make(map[coord]bool)
    dir := 0
    var loc coord
    ans := 0

    for _, obs := range obstacles {
        obsMap[coord{obs[0], obs[1]}] = true
    }

    for _, cmd := range commands {
        if cmd == -2 {
            dir = (dir + 3) % 4
        } else if cmd == -1 {
            dir = (dir + 1) % 4
        } else {
            for range cmd {
                if dir == 0 {
                    loc.y++
                    if obsMap[loc] {
                        loc.y--
                    }
                } else if dir == 1 {
                    loc.x++
                    if obsMap[loc] {
                        loc.x--
                    }
                } else if dir == 2 {
                    loc.y--
                    if obsMap[loc] {
                        loc.y++
                    }
                } else {
                    loc.x--
                    if obsMap[loc] {
                        loc.x++
                    }
                }
                ans = max(ans, loc.x * loc.x + loc.y * loc.y)
            }
        }
    }

    return ans;
}
