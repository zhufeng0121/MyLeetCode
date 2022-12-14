package Hot100

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{0, 0}
	}

	// map 中存储 index 和 数值 key 为值， value 为下标
	number := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := number[target-nums[i]]; ok {
			return []int{i, v}
		}
		number[nums[i]] = i
	}
	return []int{0, 0}
}
