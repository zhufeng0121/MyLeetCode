package Hot100

// 遍历给定数组中的元素，如果队列不为空，且当前考察元素大于队尾元素，将队尾元素移除，直到 队列为空 或者当前考察元素 小于新的队尾元素
// 当队首元素的下标小于滑动窗口左侧边界的left时，表示队首元素已经不在滑动窗口内，将其从队首移除
// 由于数组下标从0开始，因此当窗口右边界right+1大于等于窗口大小k时，意味着窗口形成。此时，队首元素就是该窗口内的最大值
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	queue := make([]int, 0) // 单调队列，存放数组元素的下标

	for i, v := range nums {
		if len(queue) == 0 || nums[queue[len(queue)-1]] >= v {
			queue = append(queue, i)
		} else if queue[len(queue)-1] < v {
			queue = queue[:len(queue)-1]
			queue = append(queue, i) // 入队
		}

		// 如果队首元素的下标 < left
		if queue[0] < i-k+1 {
			res = append(res, nums[queue[0]])
			queue = queue[1:]
		}

	}
	if len(queue) != 0 {
		res = append(res, nums[queue[0]])
		queue = queue[1:]
	}

	return res

}
