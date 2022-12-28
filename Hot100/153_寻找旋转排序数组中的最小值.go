package Hot100

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	start, end := 0, len(nums)-1

	// 终止条件是 start < end 留一个变量
	for start < end {
		mid := start + (end-start)>>1
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid + 1
		}

	}
	return nums[start]

}
