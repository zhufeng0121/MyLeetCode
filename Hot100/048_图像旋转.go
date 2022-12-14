package Hot100

/*
*  clockwise rotate 顺时针旋转
*  first reverse up to down, then swap the symmetry
*  1 2 3    7 8 9    7 4 1
*  4 5 6 => 4 5 6 => 8 5 2
*  7 8 9    1 2 3    9 6 3
 */

/*
* anticlockwise rotate 逆时针旋转
* first reverse left to right, then swap the symmetry
* 1 2 3    3 2 1    3 6 9
* 4 5 6 => 6 5 4 => 2 5 8
* 7 8 9    9 8 7    1 4 7
 */

func rotate(matrix [][]int) {
	length := len(matrix)
	// 对角线变化 矩阵 右上角的元素的性质是 纵坐标大于横坐标
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 竖直轴对称反转
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}
