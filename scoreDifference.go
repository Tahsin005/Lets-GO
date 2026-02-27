func scoreDifference(nums []int) int {
    firstPlayerActiveStatus, secondPlayerActiveStatus := true, false
    firstPlayerScore, secondPlayerScore := 0, 0

    for i, val := range nums {
        if val % 2 == 1 {
            firstPlayerActiveStatus, secondPlayerActiveStatus = secondPlayerActiveStatus, firstPlayerActiveStatus
        }

        if isEvery6thGame(i) {
            firstPlayerActiveStatus, secondPlayerActiveStatus = secondPlayerActiveStatus, firstPlayerActiveStatus
        }

        if firstPlayerActiveStatus {
            firstPlayerScore += val
        } else {
            secondPlayerScore += val
        }
    }

    return firstPlayerScore - secondPlayerScore
}

func isEvery6thGame(index int) bool {
    return (index + 1) % 6 == 0
}
