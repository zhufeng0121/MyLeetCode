package Hot100

// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make(map[int]int)
	stack := make([]int, 0, len(nums2))

	ans := make([]int, 0, len(nums1))

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			// 出栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[nums2[i]] = -1
		} else {
			res[nums2[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])

	}

	for _, v := range nums1 {
		ans = append(ans, res[v])
	}

	return ans

}
