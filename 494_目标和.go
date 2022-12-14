package main

// 先采用回溯法, 对每一个元素赋值有两种结果，- / +
func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	var backtrack func(nums []int, index, sum, target int)

	backtrack = func(nums []int, index, sum, target int) {
		if index == len(nums) {
			if sum == target {
				res++
			}
			return
		}

		// nums[i] 选择 - 号
		sum += nums[index]
		backtrack(nums, index+1, sum, target)
		sum -= nums[index]

		// nums[i] 选择 + 号
		sum -= nums[index]
		backtrack(nums, index+1, sum, target)
		sum += nums[index]

	}

	backtrack(nums, 0, 0, target)

	return res
}

// 采用DP进行优化
