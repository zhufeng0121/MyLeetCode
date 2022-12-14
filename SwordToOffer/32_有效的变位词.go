package SwordToOffer

func isAnagram(s string, t string) bool {
	// 如果 s 和 t 的长度不相等
	if len(s) != len(t) {
		return false
	}
	// 如果 s 和 t 字符串的内容相同
	if s == t {
		return false
	}
	count_map := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count_map[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if _, ok := count_map[t[i]]; !ok {
			return false

		}
		count_map[t[i]]--
	}

	for _, v := range count_map {
		if v != 0 {
			return false
		}
	}
	return true
}
