package Hot100

// 双指针，向内侧移动短板 移动短板，若遇到长板，有可能面积增大  移动长板，面积会减小
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	maxArea := 0

	for start < end {
		area := (end - start) * minUtil(height[end], height[start])
		if maxArea < area {
			maxArea = area

		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}

	}
	return maxArea

}

func minUtil(a, b int) int {
	if a < b {
		return a
	}
	return b
}
