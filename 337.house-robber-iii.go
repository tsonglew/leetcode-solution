/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rob(root *TreeNode) int {
    wr, wo := robSub(root)
    return max(wr, wo)
}

func robSub(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    withRootLeft, withoutRootLeft := robSub(root.Left)
    withRootRight, withoutRootRight := robSub(root.Right)
    return root.Val+withoutRootLeft+withoutRootRight, max(withRootLeft+withRootRight, withRootLeft+withoutRootRight, withoutRootLeft+withRootRight, withoutRootLeft+withoutRootRight) 
}

func max(nums ...int) int {
    mx := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > mx {
            mx = nums[i]
        }
    }
    return mx
}