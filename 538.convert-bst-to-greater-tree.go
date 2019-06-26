/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    if root != nil {
        convert(root, 0)
    }
    return root
}

func convert(root *TreeNode, val int) {
    if root.Right != nil {
        convert(root.Right, val)
        p := root.Right 
        for p.Left != nil {
            p = p.Left
        }
        root.Val += p.Val
    } else {
        root.Val += val
    }
    if root.Left != nil {
        convert(root.Left, root.Val)
    }
}
