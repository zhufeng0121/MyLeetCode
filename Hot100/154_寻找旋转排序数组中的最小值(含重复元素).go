package Hot100

// 寻找旋转排序数组中的最小值II

// 难点的处理还是在 nums[mid] == nums[right] 上， 需要证明 right--的处理方法不会保证将最小的元素错过

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/solutions/9474/154-find-minimum-in-rotated-sorted-array-ii-by-jyd/?orderBy=most_relevant
func findMinDuplicated(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)>>1
		if nums[mid] < nums[end] {
			// 说明最小元素的值在 start, end 之间 end是闭区间
			end = mid
		} else if nums[mid] > nums[end] {
			// 说明最小元素的值在 mid + 1，end 之间
			start = mid + 1
		} else {
			// 这块的处理要看一下
			end--
		}
	}

	return nums[start]

}
