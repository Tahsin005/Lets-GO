func mapWordWeights(words []string, weights []int) string {
    var result string
    
    for _,v := range words {
        var sum int
        for _,w := range []rune(v) {
            sum += weights[int(w - 'a')]
        }
        sum = 25 - (sum % 26)
        result += string('a' + sum)
    }

    return result
}
