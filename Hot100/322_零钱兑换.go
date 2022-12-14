package Hot100

func coinChange(coins []int, amount int) int {

	// dp 数组的定义，当目标金额为 i 时，最少需要 dp[i]个硬币凑出
	dp := make([]int, amount+1)
	// 将所有元素初始化为
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	//base case
	dp[0] = 0

	// 外层 for 循环遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = minInt(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}

}
