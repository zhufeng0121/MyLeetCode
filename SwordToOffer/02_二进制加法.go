package SwordToOffer

import "strings"

func addBinary(a string, b string) string {
	l1, l2 := len(a), len(b)

	l := l1
	if l1 > l2 {
		b = strings.Repeat("0", l1-l2) + b
	} else {
		a = strings.Repeat("0", l2-l1) + a
		l = l2
	}

	s := []byte(a)
	var carry byte = 0

	for i := l - 1; i >= 0; i-- {
		s[i] = (a[i]-'0'+b[i]-'0'+carry)%2 + '0'
		carry = (a[i] - '0' + b[i] - '0' + carry) / 2

	}
	// 不要忘记进位
	if carry == 1 {
		return "1" + string(s)
	}
	return string(s)

}
