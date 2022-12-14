package SwordToOffer

// 将该问题转换成0-1背包问题，即给定一个 可以装载 sum/2 容量的背包 和 N个物品，每个物品的重量为 nums[i].
// 是否存在一种装法，能够将背包恰好装满
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}
	n := len(nums)
	sum = sum / 2

	// 构造 dp 数组 dp[i][j] = x 表示 对于前i个物品，当前的背包容量为j时，若x为true，则说明可以恰好将背包装满，
	dp := make([][]bool, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum+1)
	}
	// base case
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}
	// 计算状态转移

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				//背包容量不足，不能装入
				dp[i][j] = dp[i-1][j]
			} else {
				// 是否装入背包，看是否存在一种情况能够恰好装满
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]

}