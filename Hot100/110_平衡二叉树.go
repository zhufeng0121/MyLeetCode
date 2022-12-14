package Hot100

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if absUtil(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)

}

func absUtil(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
