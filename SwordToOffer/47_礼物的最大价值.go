package SwordToOffer

// 非常简单地dp
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 创建二维dp数组
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	// 注意m 和 n的顺序
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = maxInt(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]

}
