package SwordToOffer

// 讲解 ：https://leetcode.cn/problems/chou-shu-lcof/solutions/182045/mian-shi-ti-49-chou-shu-dong-tai-gui-hua-qing-xi-t/
func nthUglyNumber(n int) int {
	// 使用 dp 数组来存储丑数
	dp := make([]int, n)

	// base case
	dp[0] = 1
	a, b, c := 0, 0, 0 // 下个应该通过乘2来获得新丑数的数据是第a个，同理b, c

	for i := 1; i < n; i++ {
		// 第a丑数个数需要通过乘2来得到下个丑数，第b丑数个数需要通过乘2来得到下个丑数，同理第c个数
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = minUtil(n2, minUtil(n3, n5))

		if dp[i] == n2 {
			// 第a个数已经通过乘2得到了一个新的丑数，那下个需要通过乘2得到一个新的丑数的数应该是第(a+1)个数
			a++
		}
		if dp[i] == n3 {
			// 第 b个数已经通过乘3得到了一个新的丑数，那下个需要通过乘3得到一个新的丑数的数应该是第(b+1)个数
			b++
		}
		if dp[i] == n5 {
			// 第 c个数已经通过乘5得到了一个新的丑数，那下个需要通过乘5得到一个新的丑数的数应该是第(c+1)个数
			c++
		}
	}
	return dp[n-1]

}

func minUtil(a, b int) int {
	if a < b {
		return a
	}
	return b
}
