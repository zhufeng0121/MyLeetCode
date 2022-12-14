package SwordToOffer

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		if num&1 != 0 {
			res++
		}
		num = num >> 1
	}

	return res

}
