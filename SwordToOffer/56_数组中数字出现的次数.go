package SwordToOffer

// 异或运算
func singleNumbers(nums []int) []int {
	temp, div := 0, 1
	for _, v := range nums {
		temp = temp ^ v
	}

	for temp&div == 0 {
		div <<= 1
	}

	a, b := 0, 0
	for _, v := range nums {
		if div&v != 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}

}
