package SwordToOffer

// 替换空格(这个题用go就太好写了)
func replaceSpace(s string) string {
	arr := []byte(s)

	res := make([]byte, 0)

	for _, v := range arr {
		if v != ' ' {
			res = append(res, v)
		} else {
			res = append(res, '%', '2', '0')
		}
	}
	return string(res)

}
