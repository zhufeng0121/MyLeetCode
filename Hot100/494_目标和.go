package Hot100


//TODO:待议
func findTargetSumWays(nums []int, target int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum < target || (sum+target)%2 == 1 {
		return 0
	}
	return change_v((sum+target)/2, nums)

}

func change_v(amount int, coins []int) int {
	n := len(coins)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1
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
