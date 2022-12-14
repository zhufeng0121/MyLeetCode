package SwordToOffer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])

	res := make([]int, 0)

	up, bottom := 0, m-1
	left, right := 0, n-1

	for up <= bottom && left <= right {
		// 边界条件处理
		if up == bottom {
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			up++
		} else if left == right {
			for i := up; i <= bottom; i++ {
				res = append(res, matrix[i][left])
			}
			left++
		} else {
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			up++
			for i := up; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
			for i := bottom; i >= up; i-- {
				res = append(res, matrix[i][left])
			}

		}

	}

	return res

}

//有时间试试逆序打印的写法
