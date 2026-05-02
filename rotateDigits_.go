type NUM_STATE int

const (
    NUM_GOOD NUM_STATE = 1
    NUM_VALID NUM_STATE = 2
    NUM_INVALID NUM_STATE = 3
)

func numState(n int) NUM_STATE {
    switch n {
        case 0:
            fallthrough
        case 1:
            fallthrough
        case 8:
            return NUM_VALID
        case 2:
            fallthrough 
        case 5:
            fallthrough
        case 6:
            fallthrough
        case 9:
            return NUM_GOOD
    }
    return NUM_INVALID
}

func rotateOneDigit(n int) bool {
    hasChanged := false
    for n > 0 {
        digit := n % 10
        ns := numState(digit)
        if ns == NUM_GOOD {
            hasChanged = true
        } else if ns == NUM_INVALID {
            return false
        }
        n /= 10
    }
    return hasChanged
}

func rotatedDigits(n int) int {
    cnt := 0
    for i := 1; i <= n; i ++ {
        if rotateOneDigit(i) {
            cnt++
        }
    }
    return cnt
}
