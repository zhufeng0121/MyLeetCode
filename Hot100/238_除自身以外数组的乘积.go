package Hot100

func productExceptSelf(nums []int) []int {
	// L, R 分别表示左右两侧的乘积列表
	L, R, answer := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	// L[i] 是索引 i 左侧所有元素的乘积  L[0] = 1
	L[0] = 1
	for i := 1; i < len(nums); i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	for i := 0; i < len(nums); i++ {
		answer[i] = L[i] * R[i]
	}

	return answer

}
