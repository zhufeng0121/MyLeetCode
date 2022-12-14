package Hot100

import "sort"

// 思路： 先对nums 数组进行排序，然后，采用双指针
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := [][]int{}

	// 这个 i 移动的始终是 左指针
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		// 如果当前遍历的值 > 0 说明后面相加不会出现等于0的情况，直接返回res
		if n1 > 0 {
			break
		}
		// 如果 当前数和前一个数相同，直接 continue，结果去重
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		// 左指针， 右指针
		l, r := i+1, len(nums)-1

		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res

}
