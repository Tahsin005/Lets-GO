func solveQueries(nums []int, queries []int) []int {
    valToIndices := make(map[int][]int)
    for idx, val := range nums {
        indices, ok := valToIndices[val]
        if ok {
            valToIndices[val] = append(indices, idx)
        } else {
            valToIndices[val] = []int{idx}
        }
    }

    answer := make([]int, len(queries))

    for idx, query := range queries {
        val := nums[query]
        indices := valToIndices[val]
        if len(indices) == 1 {
            answer[idx] = -1
            continue
        }

        queryPos := sort.SearchInts(indices, query)

        var leftDist int
        if queryPos == 0 {
            leftDist = (indices[0] + len(nums)) - indices[len(indices) - 1]
        } else {
            leftDist = indices[queryPos] - indices[queryPos - 1]
        }

        var rightDist int
        if queryPos == len(indices) - 1 {
            rightDist = indices[0] - (indices[len(indices) - 1] - len(nums))
        } else {
            rightDist = indices[queryPos + 1] - indices[queryPos]
        }

        answer[idx] = min(leftDist, rightDist)
    }

    return answer
}
