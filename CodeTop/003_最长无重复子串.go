package CodeTop

// 题目： 给定一个字符串s，找出不含有重复字符的最长子串的长度

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	arr := []byte(s)

	left, right := 0, 0

	res := 0

	for right < len(arr) {
		c := arr[right]
		right++

		window[c]++
		for window[c] > 1 {
			// 说明有重复元素，left需要向右进行收缩
			l := arr[left]
			left++
			window[l]--
		}

		res = max(res, right-left)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
