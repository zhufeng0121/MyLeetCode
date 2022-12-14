package Hot100

func generateParenthesisI(n int) []string {
	if n == 0 {
		return []string{}
	}

	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res

}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	// 终止条件
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}

	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}

	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

// 括号生成算法

//回溯算法的框架就是：

//result = []
//func backtrack(路劲，选择列表)
//	if 满足结束条件：
//		result.add()
//      return
//
//  for 选择 in 选择列表
//      做选择
//      backtrack(路径, 选择列表)
//      撤销选择
//
