package SwordToOffer

// 动态规划经典题目
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 的定义 以 i, j 为结尾的 最长公共子序列的长度 dp[i-1][j-1] + 1 , dp
	m, n := len(text1), len(text2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxUtil(dp[i-1][j], dp[i][j-1])

			}

		}
	}
	return dp[m][n]

}

func maxUtil(a, b int) int {
	if a > b {
		return a
	}
	return b
}
