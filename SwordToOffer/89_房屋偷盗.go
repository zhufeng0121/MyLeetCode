package SwordToOffer

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))

	// dp[i] 表示前 i 间房屋能偷窃到的最高总金额，状态转移方程
	// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	//
	dp[0] = nums[0]
	dp[1] = maxUtil(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = maxUtil(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]

}

// 空间优化版
func rob_v2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := maxUtil(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		first, second = second, maxUtil(first+nums[i], second)
	}
	return second
}
