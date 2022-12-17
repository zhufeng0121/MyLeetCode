package main

func maxAscendingSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	cur := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur += nums[i]
			res = maxInt(res, cur)
		} else {
			cur = 0
			cur += nums[i]
		}

	}
	return res

}
