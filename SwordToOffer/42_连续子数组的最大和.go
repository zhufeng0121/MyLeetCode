package SwordToOffer

//暴力法解决
func maxSubArray(nums []int) int {
	n := len(nums)
	minValue := ^(int(^uint(0) >> 1))

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			minValue = maxInt(minValue, sum)
		}
	}
	return minValue

}

//暴力法解决
func maxSubArrayI(nums []int) int {
	minValue := ^int(^uint(0) >> 1)
	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			minValue = maxInt(minValue, sum)
		}
	}
	return minValue

}

func maxSubArrayII(nums []int) int {
	minValue := ^int(^uint(0) >> 1)
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		minValue = maxInt(minValue, sum)
	}
	return minValue
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
