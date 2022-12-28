package Hot100

func longestPalindrome(s string) string {
	// 最长回文子串

	arr := []byte(s)
	// 初始值全部为 false
	dp := make([][]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]bool, len(arr))
	}

	maxLen := 0
	left, right := 0, 0

	// 注意 dp[i][j] 的含义， i <= j 因此只能填右上半部分

	// 由于 dp[i][j] 的取值取决于 dp[i+1][j-1] 因此计算的顺序应该是从下到上，从左到右

	// base case
	for i := 0; i < len(arr); i++ {
		dp[i][i] = true
	}

	for i := len(arr) - 2; i >= 0; i-- {
		for j := i + 1; j < len(arr); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			// 更新长度
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				left = i
				right = j
			}

		}
	}

	return string(arr[left : right+1])

}
