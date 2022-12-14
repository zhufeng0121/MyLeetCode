package SwordToOffer

func pivotIndex(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return -1
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	sum := 0

	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1

}
