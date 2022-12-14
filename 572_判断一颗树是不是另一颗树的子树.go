package main

func isSubtreeI(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtreeI(root.Left, subRoot) || isSubtreeI(root.Right, subRoot)
}