func maxNumberOfBalloons(text string) int {
    freq := make([]int, 26)

    for _, c := range text {
        freq[c - 'a']++
    }

    return min(
        freq['b' - 'a'],
        freq['a' - 'a'],
        freq['n' - 'a'],
        freq['l' - 'a'] / 2,
        freq['o' - 'a'] / 2,
    )
}

func min(vals ...int) int {
    res := vals[0]
    for _, v := range vals {
        if v < res {
            res = v
        }
    }
    return res
}
