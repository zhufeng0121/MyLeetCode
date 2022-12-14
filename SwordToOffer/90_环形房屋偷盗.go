package SwordToOffer

func rob_v3(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return maxUtil(nums[0], nums[1])
	}
	return maxUtil(_rob(nums[:n-1]), _rob(nums[1:]))
}

func _rob(nums []int) int {
	first, second := nums[0], maxUtil(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, maxUtil(first+v, second)
	}
	return second

}
