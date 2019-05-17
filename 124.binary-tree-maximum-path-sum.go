/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val
	pathSum(root, &maxSum)
	return maxSum
}

func pathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	pathSumLeft := pathSum(root.Left, maxSum)
	pathSumRight := pathSum(root.Right, maxSum)
	currentRootSum := pathSumLeft + pathSumRight + root.Val
	if currentRootSum > *maxSum {
		*maxSum = currentRootSum
	}
	sum := root.Val + zeroMax(pathSumLeft, pathSumRight)
	if sum > 0 {
		return sum
	}
	return 0
}

func zeroMax(a, b int) int {
	m := a
	if m < b {
		m = b
	}
	if m < 0 {
		m = 0
	}
	return m
}