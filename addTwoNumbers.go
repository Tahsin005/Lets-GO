import "math/big"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    num1 := toNumber(l1)
    num2 := toNumber(l2)

    sum := new(big.Int).Add(num1, num2)
    return toList(sum.String())
}

func toNumber(l *ListNode) *big.Int {
    str := ""
    for l != nil {
        str = string(rune('0'+l.Val)) + str
        l = l.Next
    }
    num := new(big.Int)
    num.SetString(str, 10)
    return num
}

func toList(numStr string) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    for i := len(numStr) - 1; i >= 0; i-- {
        curr.Next = &ListNode{Val: int(numStr[i] - '0')}
        curr = curr.Next
    }
    return dummy.Next
}
