package SwordToOffer

// 这个题太多的边界条件需要考虑了 例如 ""， " "， "     "，后面用这个例子 尝试写一下单测用例
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	arr := []byte(s)
	words := make([]string, 0)
	var word []byte
	for i := 0; i < len(arr); i++ {
		if arr[i] == ' ' {
			if len(word) != 0 {
				words = append(words, string(word))
				word = word[:0]
			}

		} else {
			word = append(word, arr[i])
		}
		if i == len(arr)-1 && len(word) != 0 {
			words = append(words, string(word))
		}
	}
	str := ""
	if len(words) == 0 {
		return ""
	}
	if len(words) == 1 {
		return words[0]
	}
	for i := len(words) - 1; i > 0; i-- {
		str = str + words[i] + " "
	}
	str += words[0]
	return str

}
