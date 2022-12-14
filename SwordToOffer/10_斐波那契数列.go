package SwordToOffer

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b := 0, 1
	res := 0

	for i := 1; i < n; i++ {
		res = (a + b) % (1e9 + 7)
		a = b
		b = res

	}
	return res
}