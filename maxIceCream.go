func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)

    iceCreams := 0

    for _, ice := range costs {
        if coins -= ice; coins >= 0 {
            iceCreams++
        }
    }

    return iceCreams
}
