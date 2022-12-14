package main

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	// 初始化结果矩阵
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
