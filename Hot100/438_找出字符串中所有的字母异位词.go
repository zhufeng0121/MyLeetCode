package Hot100

// 滑动窗口 + 双指针

// 1. left = right = 0
// 2. right 向右扩
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if s == "" {
		return res
	}

	need := make(map[byte]int)
	arrp := []byte(p)
	for _, v := range arrp {
		need[v]++
	}
	left, right := 0, 0
	valid := 0

	window := make(map[byte]int)

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res

}
