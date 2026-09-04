func minAddToMakeValid(s string) int {
    stack := [] byte{}
    i :=0
    for i < len(s) {
        if len(stack) > 0  &&  stack[len(stack) - 1] == '(' && s[i] == ')' {
            stack = stack[:len(stack) - 1]
        } else{
            stack = append(stack, s[i])
        }
        i++
    }

    return len(stack)
}
