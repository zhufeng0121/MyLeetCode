package interview

// 矩阵的宽度和高度相同
func rotate(matrix [][]int) {
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	times := len(matrix) / 2

	for i := 0; i < times; i++ {
		rotateLayer(matrix, top, bottom, left, right)
		top, bottom = top+1, bottom-1
		left, right = left+1, right-1
	}

}

// 一层一层的旋转，先旋转最外层，由外至内，一共 N/2 次
func rotateLayer(matrix [][]int, top, bottom, left, right int) {
	times := right - left
	for i := 0; i < times; i++ {
		temp := matrix[top][left] // 保存最坐上角的元素

		// 左边的列元素向上挪一格
		for i := top + 1; i <= bottom; i++ {
			matrix[i-1][left] = matrix[i][left]
		}

		// 底部的行元素向左挪一格
		for i := left + 1; i <= right; i++ {
			matrix[bottom][i-1] = matrix[bottom][i]
		}
		// 右边的列元素向下挪一格
		for i := bottom - 1; i >= top; i-- {
			matrix[i+1][right] = matrix[i][right]
		}
		// 上面的行元素向右挪一格
		for i := right - 1; i >= left; i-- {
			matrix[top][i+1] = matrix[top][i]
		}
		matrix[top][left+1] = temp

	}
}

// 巧妙解法

//情况一：顺时针转 90 度：先转置再左右镜像
//1 2 3               7 4 1
//4 5 6               8 5 2
//7 8 9               9 6 3
//情况二：顺时针转 180 度:先上下镜像，再左右镜像（先左右再上下也可）
//1 2 3               9 8 7
//4 5 6               6 5 4
//7 8 9               3 2 1
//情况三：顺时针转 270 度：先转置再上下镜像
//1 2 3              3 6 9
//4 5 6              2 5 8
//7 8 9              1 4 7

func rotateI(matrix [][]int) {
	m := len(matrix)

	//先转置矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//再镜像对称
	left := 0
	right := m - 1
	for left < right {
		for i := 0; i < m; i++ {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
		left++
		right--
	}
}
