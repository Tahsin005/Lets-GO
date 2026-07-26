/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func deepestLeavesSum(root *TreeNode) int {
    var maxDepth, sum int
    dfs(root, 1, &maxDepth, &sum)
    return sum
}

func dfs(node *TreeNode, depth int, maxDepth, sum *int) {
    if node == nil {
        return
    }

    if depth > *maxDepth {
        *maxDepth = depth
        *sum = node.Val
    } else if depth == *maxDepth {
        *sum += node.Val
    }

    dfs(node.Left, depth + 1, maxDepth, sum)
    dfs(node.Right, depth + 1, maxDepth, sum)
}
