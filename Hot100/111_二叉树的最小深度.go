package Hot100

// 通过 BFS来进行计算  和104 题二叉树的最大深度对比来看
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	depth := 1

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 判断是否到达终点
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 增加步数
		depth++

	}
	return depth
}
