package Hot100

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)

}
func helper(root *TreeNode, i int) int {
	if root == nil {
		return 0
	}

	temp := i*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return temp
	}

	return helper(root.Left, temp) + helper(root.Right, temp)

}




