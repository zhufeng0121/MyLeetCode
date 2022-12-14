package SwordToOffer

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
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

			if i == 0 {
				res = append(res, temp.Val)

			}
		}

	}
	return res[len(res)-1]

}

// 还有一种做法，层序遍历，先入右节点，再入左节点，最后最左最下的元素是最后一个
