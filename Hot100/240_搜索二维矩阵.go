package Hot100

func searchMatrix(matrix [][]int, target int) bool {
	// 从右上角开始搜索, m,n是下标
	m := 0
	n := len(matrix[0]) - 1

	for m < len(matrix) && n >= 0 {
		if target == matrix[m][n] {
			return true
		} else if target > matrix[m][n] {
			m++
		} else {
			n--
		}

	}
	return false

}
