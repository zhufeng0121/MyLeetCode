package SwordToOffer

// 思路是 层序遍历，返回每一层元素的最后一个即可
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			// 队头元素出队
			temp := queue[0]
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}

			if i == size-1 {
				res = append(res, temp.Val)

			}
		}

	}
	return res
}
