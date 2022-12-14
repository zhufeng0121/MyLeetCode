package Hot100

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	//对nums数组进行排序
	sort.Ints(nums)
	fmt.Println(nums)
	max := 1
	currentLen := 1

	// 未考虑重复元素的情况
	for i := 0; i < len(nums)-1; i++ {
		// 如果是重复元素，跳过
		if nums[i] == nums[i+1] {
			continue
		}
		if nums[i]+1 == nums[i+1] {
			currentLen++
			max = maxInt(currentLen, max)
		} else {
			currentLen = 1
		}
	}
	return max
}

//func maxInt(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
