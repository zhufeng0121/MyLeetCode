package Hot100

func lengthOfLongestSubstring(s string) int {
	// 用双指针试一下

	runeIndex := make(map[rune]int)

	maxLen := 0
	left := 0       // 每次计算最长子串的最左边界
	currentLen := 0 // 当前无重复子串的最大长度

	for i, v := range s {
		index, ok := runeIndex[v]
		if !ok {
			// 如果不在 map 中
			currentLen = i - left + 1
		} else {
			if index >= left {
				// 如果在 map 中, 将 left 更新到 最大Index的下一个位置
				if i == runeIndex[v]+1 {
					left = i
				} else {
					left = runeIndex[v] + 1
				}
				currentLen = i - left + 1

			} else {
				currentLen++

			}

		}
		if currentLen > maxLen {
			maxLen = currentLen
		}
		// 更新该字符串出现的最大index
		runeIndex[v] = i
	}
	return maxLen

}

// 比较工整的写法，
func lengthOfLongestSubstringI(s string) int {
	runeIndex := make(map[rune]int)

	maxLen := 0
	left := 0 // 每次计算最长子串的最左边界

	for i, v := range s {
		index, ok := runeIndex[v]
		if !ok {
			maxLen = maxInt(maxLen, i-left+1)
		} else {
			if index < left {
				maxLen = minInt(maxLen, i-left+1)
			} else {
				// left 的赋值是 index + 1
				left = index + 1
			}
		}
		runeIndex[v] = i

	}
	return maxLen

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
