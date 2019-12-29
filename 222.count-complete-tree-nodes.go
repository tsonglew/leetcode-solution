/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    q := []*TreeNode{root}
    cnt := 1
    for len(q) > 0 {
        n := q[0]
        left := n.Left
        right := n.Right
        if left != nil {
            q = append(q, left)
            cnt ++
        } else {
            return cnt
        }
        if right != nil {
            q = append(q, right)
            cnt ++
        } else {
            return cnt
        }
        q = q[1:]
    }
    return cnt
}