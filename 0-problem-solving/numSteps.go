func addOne(s []rune) string {
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == '1' {
            s[i] = '0'
        } else {
            s[i] = '1'
            break;
        }
    }

    if s[0] == '0' {
        s = append([]rune{'1'}, s...)
    }

    return string(s)
}

func numSteps(s string) (ans int) {
    for s != "1" {
        ans++

        if s[len(s) - 1] == '0' {
            s = s[:len(s) - 1]
        } else {
            s = addOne([]rune(s))
        }
    }

    return
}