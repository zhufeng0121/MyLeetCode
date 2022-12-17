package Hot100

import "sort"

// 会超时，不过大致思路是这样的

// 先对宽度w进行升序排序，如果遇到w相同的情况，则按照高度h进行降序。之后按照h 计算LIS
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)

	// 宽度升序，宽度相同按照高度降序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})

	// 对高度数组进行计算lengthLIS

	height := make([]int, n)

	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}

	return lengthOfLIS(height)

}
