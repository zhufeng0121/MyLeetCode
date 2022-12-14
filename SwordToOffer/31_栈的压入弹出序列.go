package SwordToOffer

// 题解： https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/solutions/1785641/zhan-de-ya-ru-dan-chu-xu-lie-by-leetcode-6myl/
func validateStackSequences(pushed []int, popped []int) bool {

	stack := make([]int, 0)

	index := 0

	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false

}
