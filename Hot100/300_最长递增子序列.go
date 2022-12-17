package Hot100

// dp[i] 定义 以 nums[i]结尾的最长递增子序列的长度，初始值为1
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = maxInt(res, dp[i])
	}
	return res
}
