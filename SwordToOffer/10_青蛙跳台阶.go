package SwordToOffer

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	a, b := 1, 1
	res := 0

	for i := 1; i < n; i++ {
		res = (a + b) % (1e9 + 7)
		a = b
		b = res

	}
	return res

}
