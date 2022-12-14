package SwordToOffer

// 思路 记录 两个值
// 1. 数组中的某个元素
// 2. 次数， 如果下一个元素和记录中的元素值相同，则 times++ 否则,减少 times ，如果 times 为 0，将元素设置为下一个，times == 1
func majorityElement(nums []int) int {

	times := 1
	cur := nums[0]

	for i := 1; i < len(nums); i++ {
		if times == 0 {
			cur = nums[i]
			times++
			continue
		}
		if nums[i] == cur {
			times++
		} else {
			times--
		}
	}
	return cur

}
