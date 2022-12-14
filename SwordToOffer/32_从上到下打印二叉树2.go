package SwordToOffer

// 通过记录 当前队列的长度，可以获取每一层的元素个数
func levelOrderI(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := make([]int, 0)
		size := len(queue)

		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
			level = append(level, temp.Val)

		}
		res = append(res, level)

	}
	return res

}
