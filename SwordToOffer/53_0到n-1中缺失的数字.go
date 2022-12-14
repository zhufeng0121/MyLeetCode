package SwordToOffer

// 思路，应该是二分查找，可以利用这个性质，在缺失的数字之前，i == nums[i], 在缺失的之后数字以及后面， nums[i] < i

func missingNumber(nums []int) int {
	start, end := 0, len(nums)-1

	// 不用担心 会遗漏 循环的退出条件是 left = right + 1， 因此即使mid 恰好是 缺失数组的第一个
	// 最后 left  = right + 1 也会补救回来
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] == mid {
			start = mid + 1
		} else {
			end = mid - 1

		}
	}
	return start

}
