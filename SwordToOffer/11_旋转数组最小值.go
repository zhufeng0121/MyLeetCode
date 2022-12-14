package SwordToOffer

// 题解
// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solutions/340801/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-by-leetcode-s/
func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		mid := left + (right-left)>>1
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
			//无论 numbers[high] 是不是最小值，都有一个它的numbers[pivot]，可以忽略二分查找区间的右端点
		} else {
			right--

		}

	}
	return numbers[left]

}
