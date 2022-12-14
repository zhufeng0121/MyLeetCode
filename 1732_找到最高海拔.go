package main

// 前缀和
func largestAltitude(gain []int) int {
	maxLevel := 0

	prefixSumArr := make([]int, len(gain)+1)
	prefixSumArr[0] = 0

	for i := 0; i < len(gain); i++ {
		prefixSumArr[i+1] = prefixSumArr[i] + gain[i]
		if prefixSumArr[i+1] > maxLevel {
			maxLevel = prefixSumArr[i+1]
		}
	}
	return maxLevel

}
