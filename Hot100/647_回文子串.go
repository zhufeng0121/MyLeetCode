package Hot100

func countSubstrings(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := 0

	// i <= j
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					res++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res

}
