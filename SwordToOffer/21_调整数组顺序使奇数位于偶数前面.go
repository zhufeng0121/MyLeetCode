package SwordToOffer

// 双指针
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		for left < right && !isOdd(nums[right]) {
			right--
		}
		for left < right && isOdd(nums[left]) {
			left++
		}

		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums

}

func isOdd(item int) bool {
	if item&1 == 1 {
		return true
	}
	return false
}