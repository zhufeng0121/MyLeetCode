package Hot100

// 采用单调栈解决
func trap(height []int) int {
	// 单调栈，存放序号
	stack := make([]int, 0)
	sum := 0

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			h := height[stack[len(stack)-1]] // 将要出栈的值记下来
			stack = stack[:len(stack)-1]     // 出栈
			if len(stack) == 0 {
				// 如果栈空了，跳出
				break
			}
			distance := i - stack[len(stack)-1] - 1 // 横向距离
			min := minInt(height[stack[len(stack)-1]], height[i])
			sum += distance * (min - h)

		}
		stack = append(stack, i)
	}
	return sum

}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
