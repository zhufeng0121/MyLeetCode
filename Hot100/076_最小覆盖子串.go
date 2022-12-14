package Hot100

// 思路 : 滑动窗口 双指针
//  1. 初始化 left, right = 0,索引 左开右闭 区间 称之为一个窗口
//  2. 不断 向右扩，直到 窗口中的子串 涵盖 t中的所有字符
//  3. left 窗口开始缩小，直到窗口中的字符串不在符合要求
//  4. 重复第2步和第3步，直到right到达字符串的尽头
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	arrt := []byte(t)
	need := make(map[byte]int, 0)
	for _, v := range arrt {
		need[v]++
	}

	left, right := 0, 0
	valid := 0

	//  记录最小子串的索引及长度
	start, length := 0, int(^uint(0)>>1)

	window := make(map[byte]int, 0)

	for right < len(s) {
		// 2. 向右扩，直至涵盖所有字符串
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 3. 满足包含条件，缩小left
		for valid == len(need) {
			//更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}

			c := s[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}

		}

	}

	arrs := []byte(s)
	// 这块有坑
	if length == int(^uint(0)>>1) {
		return ""
	}
	return string(arrs[start : start+length])

}
