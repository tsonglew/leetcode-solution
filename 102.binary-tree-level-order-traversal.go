/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	results := [][]int{}
	nodes := []*TreeNode{root}
	newNodes := []*TreeNode{}
	for len(nodes) != 0 {
		currentVal := []int{}
		for i := range nodes {
			if nodes[i] != nil {
				newNodes = append(newNodes, nodes[i].Left, nodes[i].Right)
				currentVal = append(currentVal, nodes[i].Val)
			}
		}
		nodes = newNodes
		newNodes = []*TreeNode{}
		if len(currentVal) > 0 {
			results = append(results, append([]int{}, currentVal...))
		}
	}
	return results
}