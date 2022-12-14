package Hot100

// 先用双重for循环来做一下(时间超限)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
			res[i] = 0
		}
	}
	return res

}

// 采用单调栈来解决
func dailyTemperaturesI(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}

	for i := 0; i < length; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex

		}

		stack = append(stack, i)
	}
	return ans

}
