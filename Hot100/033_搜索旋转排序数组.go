package Hot100

// 第一种写法
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[0] {
			if nums[mid] > target && target >= nums[0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1

}

// 第二种写法
func searchI(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1

}
