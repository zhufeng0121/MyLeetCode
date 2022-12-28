package Hot100

// 利用 栈 判断括号的有效性
func isValid(s string) bool {
	stack := make([]byte, 0)
	arr := []byte(s)
	for _, v := range arr {
		if isLeftPart(v) {
			stack = append(stack, v)
		} else {
			if len(stack) != 0 && leftOf(v) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isLeftPart(b byte) bool {
	if b == '(' || b == '[' || b == '{' {
		return true
	}
	return false
}

func leftOf(b byte) byte {
	if b == ')' {
		return '('
	}
	if b == '}' {
		return '{'
	}
	return '['
}
