/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // method 1
 func countNodes1(root *TreeNode) int {
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

// method 2
func countNodes2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1+countNodes(root.Left)+countNodes(root.Right)
}