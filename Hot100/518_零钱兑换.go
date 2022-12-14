package Hot100

// 动态规划 dp[i][j] 若只使用前 i 个物品，当背包容量为 j 时，dp[i][j] 种方法可以装满背包
func change(amount int, coins []int) int {
	n := len(coins)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}

	// 如果目标金额是0，默认有一种凑法
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	// 如果不使用任何面值，就无法凑出任何金额
	for i := 0; i <= amount; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]

			}
		}
	}

	return dp[n][amount]

}
