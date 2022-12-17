package Hot100

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] 以 nums[i]结尾的
	dp := make([]int, len(nums))
	// base case
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1]+nums[i], nums[i])
	}

	res := math.MinInt
	for i := 0; i < len(dp); i++ {
		res = maxInt(res, dp[i])
	}
	return res
}

// dp[i] 的结果计算只依赖于 dp[i-1] 因此可以进行状态压缩

func maxSubArrayCompress(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp_0, dp_1 := nums[0], 0

	res := dp_0

	for i := 1; i < len(nums); i++ {
		dp_1 = maxInt(nums[i], nums[i]+dp_0)
		dp_0 = dp_1

		res = maxInt(res, dp_0)
	}
	return res

}
