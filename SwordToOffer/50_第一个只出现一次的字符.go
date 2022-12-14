package SwordToOffer

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}

	arr := []byte(s)
	countMap := make(map[byte]int)

	for _, v := range arr {
		countMap[v]++
	}

	for _, v := range arr {
		if countMap[v] == 1 {
			return v
		}
	}
	return ' '

}