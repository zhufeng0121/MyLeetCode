package SwordToOffer

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)

	// 根节点先入队
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		// 一定要搞清楚 这个 min 和max 是怎么算的
		min := ^(int(^uint(0) >> 1))
		for i := 0; i < size; i++ {
			// 出队
			temp := queue[0]
			queue = queue[1:]
			min = maxUtil(min, temp.Val)

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}

		}
		res = append(res, min)
	}
	return res

}
