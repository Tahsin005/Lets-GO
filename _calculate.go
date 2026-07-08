func calculate(s string) int {
    stack := []int{}
    num := 0
    op := '+'

    for i, c := range s {
        if c >= '0' && c <= '9' {
            num = num * 10 + int(c - '0')
        }

        if c == '/' || c == '*' || c == '+' || c == '-' || i == len(s) - 1 {
            switch op {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                tmp := stack[len(stack) - 1] * num
                stack[len(stack) - 1] = tmp
            case '/':
                tmp := stack[len(stack) - 1] / num
                stack[len(stack) - 1] = tmp
             }

            op = c
            num = 0
        } 
    }

    sum := 0
    for _, v := range stack {
        sum += v
    }

    return sum
}
