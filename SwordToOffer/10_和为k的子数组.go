package SwordToOffer

// 可以利用前缀和来实现
func subarraySum(nums []int, k int) int {
	presum := make([]int, len(nums)+1)
	presum[0] = 0

	// 前缀数组的含义 : preSum[i] 就是 nums[0 .. i-1]的和
	// 想求 nums[i .. j]的和， 可以采用 preSum[j+1] -preSum[i]
	for i := 0; i < len(nums); i++ {
		presum[i+1] = presum[i] + nums[i]
	}

	res := 0

	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if presum[i]-presum[j] == k {
				res++
			}
		}
	}
	return res

}
