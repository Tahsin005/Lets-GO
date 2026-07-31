func minimumPushes(word string) int {
    freq := make(map[rune]int)
    for _, v := range word {
        freq[v]++
    }

    counts := make([]int, 0, len(freq))
    for _, c := range freq {
        counts = append(counts, c)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(counts)))

    result := 0
    for i, c := range counts {
        pushes := i / 8 + 1
        result += pushes * c
    }

    return result
}
