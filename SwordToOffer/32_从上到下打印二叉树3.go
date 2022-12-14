package SwordToOffer

func levelOrderIII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := make([]*TreeNode, 0)
	level := 0
	queue = append(queue, root)

	for len(queue) > 0 {
		arr := []int{}
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
			arr = append(arr, temp.Val)

		}

		if level%2 == 0 {
			res = append(res, arr)
		} else {
			reverse := make([]int, 0, len(arr))
			for i := len(arr) - 1; i >= 0; i-- {
				reverse = append(reverse, arr[i])
			}
			res = append(res, reverse)
		}
		level++
	}
	return res

}
