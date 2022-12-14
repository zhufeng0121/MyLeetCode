package SwordToOffer

// 双指针
// 同: 剑指 Offer II 006. 排序数组中两个数字之和
func twoSum(nums []int, target int) []int {

	if len(nums) == 0 || len(nums) == 1 {
		return []int{}
	}

	start, end := 0, len(nums)-1

	for start < end {
		if nums[start]+nums[end] == target {
			return []int{nums[start], nums[end]}
		} else if nums[start]+nums[end] > target {
			end--
		} else if nums[start]+nums[end] < target {
			start++
		}
	}
	return []int{}

}
