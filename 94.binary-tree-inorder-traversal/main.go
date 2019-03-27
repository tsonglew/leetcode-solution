package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := &[]int{}

	if root != nil {
		if root.Left != nil {
			travelSub(root.Left, result)
		}
		(*result) = append(*result, root.Val)
		if root.Right != nil {
			travelSub(root.Right, result)
		}
	}
	return *result
}

func travelSub(root *TreeNode, result *[]int) {
	if root.Left != nil {
		travelSub(root.Left, result)
	}
	(*result) = append(*result, root.Val)
	if root.Right != nil {
		travelSub(root.Right, result)
	}
}
