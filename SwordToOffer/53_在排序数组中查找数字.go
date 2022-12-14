package SwordToOffer

func search(nums []int, target int) int {
	left := findLeft(nums, target)
	right := findRight(nums, target)

	if left == -1 || right == -1 {
		return 0
	}
	return findRight(nums, target) - findLeft(nums, target) + 1

}

// 二分查找 左边界
func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}

	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left

}

// 二分查找 右边界
func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}
	return right

}
