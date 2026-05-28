func findArray(pref []int) []int {
    returnArr := make([]int, len(pref))
    runningXOR := 0

    for i := 0; i < len(pref); i++ {
        returnArr[i] = runningXOR ^ pref[i]
        runningXOR = runningXOR ^ returnArr[i]
    }

    return returnArr
}
