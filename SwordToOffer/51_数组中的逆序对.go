package SwordToOffer

func reversePairs(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	vice := make([]int, len(nums))

	for i := 0; i < len(vice); i++ {
		vice[i] = nums[i]
	}

	return reversePairsCore(nums, vice, 0, len(nums)-1)

}

func reversePairsCore(nums []int, vice []int, start, end int) int {
	if start == end {
		vice[start] = nums[start]
		return 0
	}

	length := (end - start) / 2

	left := reversePairsCore()
	right := reversePairsCore()

	i := start + length

}
