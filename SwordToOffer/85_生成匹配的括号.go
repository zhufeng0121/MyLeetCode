package SwordToOffer

// 生成匹配的括号
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]string, 0)
	track := make([]byte, 0)

	// 将 left, right 初始化为 n, n
	backtrack(n, n, &track, &res)
	return res

}

// left, right 分别为 可用的左括号和右括号数量
func backtrack(left, right int, track *[]byte, res *[]string) {

	//进行剪枝
	if left < 0 || right < 0 {
		return
	}
	// 如果左括号剩下的多，说明不合法
	if right < left {
		return
	}
	// 当所有括号都恰好用完时，得到一个合法的括号组合
	if left == 0 && right == 0 {
		*res = append(*res, string(*track))
		return
	}

	//尝试添加一个左括号
	*track = append(*track, '(')
	backtrack(left-1, right, track, res)
	*track = (*track)[:len(*track)-1]

	// 尝试添加一个右括号
	*track = append(*track, ')')
	backtrack(left, right-1, track, res)
	*track = (*track)[:len(*track)-1]

}
