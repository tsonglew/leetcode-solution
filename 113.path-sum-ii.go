/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) [][]int {
    result := [][]int{}
    pathSumSub(root, sum, []int{}, &result)
    return result
}

func pathSumSub(root *TreeNode, sum int, current []int, result *[][]int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil && root.Val == sum {
        *result = append(*result, append(append([]int{}, current...), root.Val))
    }
    pathSumSub(root.Left, sum - root.Val, append(current, root.Val), result)
    pathSumSub(root.Right, sum - root.Val, append(current, root.Val), result)
}
