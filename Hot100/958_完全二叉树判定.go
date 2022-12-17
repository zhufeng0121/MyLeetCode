package Hot100

// 层序遍历 找到第一个空节点后，后面如果出现非空节点，则说明不是完全二叉树
func isCompleteTree(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 && q[0] != nil {
		q = append(q, q[0].Left, q[0].Right)
		q = q[1:]
	}

	for _, v := range q {
		if v != nil {
			return false
		}
	}

	return true
}
