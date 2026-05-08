func deckRevealedIncreasing(deck []int) []int {
    ans := make([]int, len(deck))
    slices.Sort(deck)

    ans[0] = deck[0]
    x := 0
    for i := 1; i < len(deck); i++ {
        for range 2 {
            x++
            if x == len(ans) {
                x = 0
            }
            for ans[x] > 0 {
                x++
                if x == len(ans) {
                    x = 0
                }
            }
        }
        ans[x] = deck[i]
    }

    return ans
}