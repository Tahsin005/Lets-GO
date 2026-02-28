const mod = 1e9 + 7
var memo [1e5 + 1]int

func init() {
    var x int
    for i := range memo {
        x <<= bits.Len(uint(i))
        x |= i
        x %= mod
        memo[i] = x
    }
}

func concatenatedBinary(n int) int {
    return memo[n]
}
