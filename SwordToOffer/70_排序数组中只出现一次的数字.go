package SwordToOffer

func singleNonDuplicate(nums []int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)>>1

		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return mid
		}
	}

}
