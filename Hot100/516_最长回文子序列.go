package Hot100

// dp[i][j] 再子串s[i..j]中，最长回文子序列的长度是 dp[i][j]

// 如果 s[i] == s[j] dp[i][j] = dp[i+1][j-1] + 2
// 否则 dp[i][j] = max(dp[i+1][j], dp[i][j-1])
// 要求的就是dp[0][n-1] 整个s的最长回文子序列的长度

func longestPalindromeSubseq(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	//base case
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// i 肯定小于 j 对于 i > j 的位置 不存在子序列，初始化为0
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 状态转移方程
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]

}
