package Hot100

// 思路: 依然是 滑动窗口 和 双指针

// 1. 双指针，初始化 left = right = 0, 窗口为 [left, right)
// 2. right 向右扩，如果满足条件，停止, 返回true，直到window中的长度大于need
// 3. left  向左走，直至不满足条件
// 4. 重复 步骤 2 和 3
func checkInclusion(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" {
		return true
	}
	need := make(map[byte]int)
	arr_s2 := []byte(s2)
	for _, v := range arr_s2 {
		need[v]++
	}
	window := make(map[byte]int)

	left, right := 0, 0
	valid := 0

	for right < len(s1) {
		c := s1[right]
		right++
		if _, ok := need[c]; ok {
			if need[c] == window[c] {
				valid++
			}
			window[c]++
		}

		// 判断是否要收缩 (窗口长度相等吗，同时 valid相等)
		for (right - left) > len(s2) {
			if valid == len(need) {
				return true
			}
			d := s1[left]
			left++
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}

		}

	}
	return false

}
