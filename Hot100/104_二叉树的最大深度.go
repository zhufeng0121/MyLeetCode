package Hot100

// 二叉树的最大深度 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))

}
