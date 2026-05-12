func minimumEffort(tasks [][]int) int {
    sort.Slice(tasks, func(i, j int) bool {
        diffI := tasks[i][1] - tasks[i][0]
        diffJ := tasks[j][1] - tasks[j][0]
        return diffI > diffJ
    })
    currentEnergy := 0
    initialEnergy := 0
    for _, task := range tasks {
        actual := task[0]
        minimum := task[1]
        if currentEnergy < minimum {
            initialEnergy += (minimum - currentEnergy)
            currentEnergy = minimum
        }
        currentEnergy -= actual
    }
    return initialEnergy
}
